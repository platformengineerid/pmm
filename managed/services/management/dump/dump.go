// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package dump exposes PMM Dump API.
package dump

import (
	"bytes"
	"context"
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/jlaffaye/ftp"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/reform.v1"

	dumpv1beta1 "github.com/percona/pmm/api/managementpb/dump"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/services/dump"
	"github.com/percona/pmm/managed/services/grafana"
)

type Service struct {
	db *reform.DB
	l  *logrus.Entry

	dumpService   dumpService
	grafanaClient *grafana.Client

	dumpv1beta1.UnimplementedDumpsServer
}

func New(db *reform.DB, grafanaClient *grafana.Client, dumpService dumpService) *Service {
	return &Service{
		db:            db,
		dumpService:   dumpService,
		grafanaClient: grafanaClient,
		l:             logrus.WithField("component", "management/dump"),
	}
}

func (s *Service) StartDump(ctx context.Context, req *dumpv1beta1.StartDumpRequest) (*dumpv1beta1.StartDumpResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("can't get request metadata")
	}

	// Here we're trying to extract authentication credentials from incoming request. We need to forward them to pmm-dump tool.
	authHeader, cookieHeader := md.Get("grpcgateway-authorization"), md.Get("grpcgateway-cookie")

	// pmm-dump supports user/pass authentication, API token or cookie.
	var token, cookie, user, password string
	if len(authHeader) != 0 {
		// If auth header type is `Basic` try to extract user and password.
		if basic, ok := strings.CutPrefix(authHeader[0], "Basic"); ok {
			decodedBasic, err := base64.StdEncoding.DecodeString(strings.TrimSpace(basic))
			if err != nil {
				return nil, errors.Wrap(err, "failed to decode basic authorization header")
			}

			s := strings.Split(string(decodedBasic), ":")
			if len(s) < 2 {
				return nil, errors.New("failed to parse basic authorization header")
			}
			user, password = s[0], s[1]
		}

		// If auth header type is `Basic` try to extract token.
		if bearer, ok := strings.CutPrefix(authHeader[0], "Bearer"); ok {
			token = strings.TrimSpace(bearer)
		}
	}

	// If auth cookie is present try to extract cookie value.
	if len(cookieHeader) != 0 {
		cookies := strings.Split(cookieHeader[0], ";")
		for _, c := range cookies {
			if auth, ok := strings.CutPrefix(strings.TrimSpace(c), "grafana_session="); ok {
				cookie = auth
			}
		}
	}

	params := &dump.Params{
		APIKey:       token,
		Cookie:       cookie,
		User:         user,
		Password:     password,
		ServiceNames: req.ServiceNames,
		ExportQAN:    req.ExportQan,
		IgnoreLoad:   req.IgnoreLoad,
	}

	if req.StartTime != nil {
		params.StartTime = req.StartTime.AsTime()
	}

	if req.EndTime != nil {
		params.EndTime = req.EndTime.AsTime()
	}

	dumpID, err := s.dumpService.StartDump(params)
	if err != nil {
		return nil, err
	}

	return &dumpv1beta1.StartDumpResponse{DumpId: dumpID}, nil
}

func (s *Service) ListDumps(_ context.Context, _ *dumpv1beta1.ListDumpsRequest) (*dumpv1beta1.ListDumpsResponse, error) {
	dumps, err := models.FindDumps(s.db.Querier, models.DumpFilters{})
	if err != nil {
		return nil, err
	}

	dumpsResponse := make([]*dumpv1beta1.Dump, 0, len(dumps))
	for _, dump := range dumps {
		d, err := convertDump(dump)
		if err != nil {
			return nil, err
		}

		dumpsResponse = append(dumpsResponse, d)
	}

	return &dumpv1beta1.ListDumpsResponse{
		Dumps: dumpsResponse,
	}, nil
}

func (s *Service) DeleteDump(_ context.Context, req *dumpv1beta1.DeleteDumpRequest) (*dumpv1beta1.DeleteDumpResponse, error) {
	for _, id := range req.DumpIds {
		if err := s.dumpService.DeleteDump(id); err != nil {
			return nil, err
		}
	}

	return &dumpv1beta1.DeleteDumpResponse{}, nil
}

func (s *Service) GetDumpLogs(_ context.Context, req *dumpv1beta1.GetLogsRequest) (*dumpv1beta1.GetLogsResponse, error) {
	filter := models.DumpLogsFilter{
		DumpID: req.DumpId,
		Offset: int(req.Offset),
	}
	if req.Limit > 0 {
		filter.Limit = pointer.ToInt(int(req.Limit))
	}

	dumpLogs, err := models.FindDumpLogs(s.db.Querier, filter)
	if err != nil {
		return nil, err
	}

	res := &dumpv1beta1.GetLogsResponse{
		Logs: make([]*dumpv1beta1.LogChunk, 0, len(dumpLogs)),
	}
	for _, log := range dumpLogs {
		if log.LastChunk {
			res.End = true
			break
		}
		res.Logs = append(res.Logs, &dumpv1beta1.LogChunk{
			ChunkId: log.ChunkID,
			Data:    log.Data,
		})
	}

	return res, nil
}

func (s *Service) UploadDump(_ context.Context, req *dumpv1beta1.UploadDumpRequest) (*dumpv1beta1.UploadDumpResponse, error) {
	files, err := s.dumpService.GetFilePathsForDumps(req.DumpIds)
	if err != nil {
		return nil, err
	}

	conn, err := ftp.Dial(req.FtpParameters.GetAddress(), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to dial remote FTP server: %v", err)
	}

	err = conn.Login(req.FtpParameters.User, req.FtpParameters.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Failed to authenticate on remote FTP server")
	}

	for _, f := range files {
		fileName := filepath.Base(f)
		data, err := os.ReadFile(f) //nolint:gosec
		if err != nil {
			return nil, errors.Wrap(err, "failed to read dump file")
		}

		err = conn.Stor(fileName, bytes.NewReader(data))
		if err != nil {
			return nil, errors.Wrap(err, "failed to write dump file on FTP server")
		}
	}

	if err := conn.Quit(); err != nil {
		return nil, err
	}

	return &dumpv1beta1.UploadDumpResponse{}, nil
}

func convertDump(dump *models.Dump) (*dumpv1beta1.Dump, error) {
	ds, err := convertDumpStatus(dump.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert dump ds")
	}

	return &dumpv1beta1.Dump{
		DumpId:       dump.ID,
		Status:       ds,
		ServiceNames: dump.ServiceNames,
		StartTime:    timestamppb.New(dump.StartTime),
		EndTime:      timestamppb.New(dump.EndTime),
		CreatedAt:    timestamppb.New(dump.CreatedAt),
	}, nil
}

func convertDumpStatus(status models.DumpStatus) (dumpv1beta1.DumpStatus, error) {
	switch status {
	case models.DumpStatusSuccess:
		return dumpv1beta1.DumpStatus_BACKUP_STATUS_SUCCESS, nil
	case models.DumpStatusError:
		return dumpv1beta1.DumpStatus_BACKUP_STATUS_ERROR, nil
	case models.DumpStatusInProgress:
		return dumpv1beta1.DumpStatus_BACKUP_STATUS_IN_PROGRESS, nil
	default:
		return dumpv1beta1.DumpStatus_BACKUP_STATUS_INVALID, errors.Errorf("invalid status '%s'", status)
	}
}