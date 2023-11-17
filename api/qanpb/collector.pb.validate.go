// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: qanpb/collector.proto

package qanv1beta1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	inventorypbv1 "github.com/percona/pmm/api/inventorypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = inventorypbv1.AgentType(0)
)

// Validate checks the field values on CollectRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CollectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CollectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CollectRequestMultiError,
// or nil if none found.
func (m *CollectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CollectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMetricsBucket() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CollectRequestValidationError{
						field:  fmt.Sprintf("MetricsBucket[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CollectRequestValidationError{
						field:  fmt.Sprintf("MetricsBucket[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CollectRequestValidationError{
					field:  fmt.Sprintf("MetricsBucket[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CollectRequestMultiError(errors)
	}

	return nil
}

// CollectRequestMultiError is an error wrapping multiple validation errors
// returned by CollectRequest.ValidateAll() if the designated constraints
// aren't met.
type CollectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CollectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CollectRequestMultiError) AllErrors() []error { return m }

// CollectRequestValidationError is the validation error returned by
// CollectRequest.Validate if the designated constraints aren't met.
type CollectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectRequestValidationError) ErrorName() string { return "CollectRequestValidationError" }

// Error satisfies the builtin error interface
func (e CollectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectRequestValidationError{}

// Validate checks the field values on MetricsBucket with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MetricsBucket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsBucket with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetricsBucketMultiError, or
// nil if none found.
func (m *MetricsBucket) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsBucket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Queryid

	// no validation rules for ExplainFingerprint

	// no validation rules for PlaceholdersCount

	// no validation rules for Fingerprint

	// no validation rules for ServiceName

	// no validation rules for Database

	// no validation rules for Schema

	// no validation rules for Username

	// no validation rules for ClientHost

	// no validation rules for NodeId

	// no validation rules for NodeName

	// no validation rules for NodeType

	// no validation rules for MachineId

	// no validation rules for ContainerId

	// no validation rules for ContainerName

	// no validation rules for NodeModel

	// no validation rules for Region

	// no validation rules for Az

	// no validation rules for ServiceId

	// no validation rules for ServiceType

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	// no validation rules for AgentId

	// no validation rules for AgentType

	// no validation rules for Labels

	// no validation rules for PeriodStartUnixSecs

	// no validation rules for PeriodLengthSecs

	// no validation rules for Example

	// no validation rules for ExampleFormat

	// no validation rules for IsTruncated

	// no validation rules for ExampleType

	// no validation rules for ExampleMetrics

	// no validation rules for NumQueriesWithWarnings

	// no validation rules for Warnings

	// no validation rules for NumQueriesWithErrors

	// no validation rules for Errors

	// no validation rules for NumQueries

	// no validation rules for MQueryTimeCnt

	// no validation rules for MQueryTimeSum

	// no validation rules for MQueryTimeMin

	// no validation rules for MQueryTimeMax

	// no validation rules for MQueryTimeP99

	// no validation rules for MLockTimeCnt

	// no validation rules for MLockTimeSum

	// no validation rules for MLockTimeMin

	// no validation rules for MLockTimeMax

	// no validation rules for MLockTimeP99

	// no validation rules for MRowsSentCnt

	// no validation rules for MRowsSentSum

	// no validation rules for MRowsSentMin

	// no validation rules for MRowsSentMax

	// no validation rules for MRowsSentP99

	// no validation rules for MRowsExaminedCnt

	// no validation rules for MRowsExaminedSum

	// no validation rules for MRowsExaminedMin

	// no validation rules for MRowsExaminedMax

	// no validation rules for MRowsExaminedP99

	// no validation rules for MRowsAffectedCnt

	// no validation rules for MRowsAffectedSum

	// no validation rules for MRowsAffectedMin

	// no validation rules for MRowsAffectedMax

	// no validation rules for MRowsAffectedP99

	// no validation rules for MRowsReadCnt

	// no validation rules for MRowsReadSum

	// no validation rules for MRowsReadMin

	// no validation rules for MRowsReadMax

	// no validation rules for MRowsReadP99

	// no validation rules for MMergePassesCnt

	// no validation rules for MMergePassesSum

	// no validation rules for MMergePassesMin

	// no validation rules for MMergePassesMax

	// no validation rules for MMergePassesP99

	// no validation rules for MInnodbIoROpsCnt

	// no validation rules for MInnodbIoROpsSum

	// no validation rules for MInnodbIoROpsMin

	// no validation rules for MInnodbIoROpsMax

	// no validation rules for MInnodbIoROpsP99

	// no validation rules for MInnodbIoRBytesCnt

	// no validation rules for MInnodbIoRBytesSum

	// no validation rules for MInnodbIoRBytesMin

	// no validation rules for MInnodbIoRBytesMax

	// no validation rules for MInnodbIoRBytesP99

	// no validation rules for MInnodbIoRWaitCnt

	// no validation rules for MInnodbIoRWaitSum

	// no validation rules for MInnodbIoRWaitMin

	// no validation rules for MInnodbIoRWaitMax

	// no validation rules for MInnodbIoRWaitP99

	// no validation rules for MInnodbRecLockWaitCnt

	// no validation rules for MInnodbRecLockWaitSum

	// no validation rules for MInnodbRecLockWaitMin

	// no validation rules for MInnodbRecLockWaitMax

	// no validation rules for MInnodbRecLockWaitP99

	// no validation rules for MInnodbQueueWaitCnt

	// no validation rules for MInnodbQueueWaitSum

	// no validation rules for MInnodbQueueWaitMin

	// no validation rules for MInnodbQueueWaitMax

	// no validation rules for MInnodbQueueWaitP99

	// no validation rules for MInnodbPagesDistinctCnt

	// no validation rules for MInnodbPagesDistinctSum

	// no validation rules for MInnodbPagesDistinctMin

	// no validation rules for MInnodbPagesDistinctMax

	// no validation rules for MInnodbPagesDistinctP99

	// no validation rules for MQueryLengthCnt

	// no validation rules for MQueryLengthSum

	// no validation rules for MQueryLengthMin

	// no validation rules for MQueryLengthMax

	// no validation rules for MQueryLengthP99

	// no validation rules for MBytesSentCnt

	// no validation rules for MBytesSentSum

	// no validation rules for MBytesSentMin

	// no validation rules for MBytesSentMax

	// no validation rules for MBytesSentP99

	// no validation rules for MTmpTablesCnt

	// no validation rules for MTmpTablesSum

	// no validation rules for MTmpTablesMin

	// no validation rules for MTmpTablesMax

	// no validation rules for MTmpTablesP99

	// no validation rules for MTmpDiskTablesCnt

	// no validation rules for MTmpDiskTablesSum

	// no validation rules for MTmpDiskTablesMin

	// no validation rules for MTmpDiskTablesMax

	// no validation rules for MTmpDiskTablesP99

	// no validation rules for MTmpTableSizesCnt

	// no validation rules for MTmpTableSizesSum

	// no validation rules for MTmpTableSizesMin

	// no validation rules for MTmpTableSizesMax

	// no validation rules for MTmpTableSizesP99

	// no validation rules for MQcHitCnt

	// no validation rules for MQcHitSum

	// no validation rules for MFullScanCnt

	// no validation rules for MFullScanSum

	// no validation rules for MFullJoinCnt

	// no validation rules for MFullJoinSum

	// no validation rules for MTmpTableCnt

	// no validation rules for MTmpTableSum

	// no validation rules for MTmpTableOnDiskCnt

	// no validation rules for MTmpTableOnDiskSum

	// no validation rules for MFilesortCnt

	// no validation rules for MFilesortSum

	// no validation rules for MFilesortOnDiskCnt

	// no validation rules for MFilesortOnDiskSum

	// no validation rules for MSelectFullRangeJoinCnt

	// no validation rules for MSelectFullRangeJoinSum

	// no validation rules for MSelectRangeCnt

	// no validation rules for MSelectRangeSum

	// no validation rules for MSelectRangeCheckCnt

	// no validation rules for MSelectRangeCheckSum

	// no validation rules for MSortRangeCnt

	// no validation rules for MSortRangeSum

	// no validation rules for MSortRowsCnt

	// no validation rules for MSortRowsSum

	// no validation rules for MSortScanCnt

	// no validation rules for MSortScanSum

	// no validation rules for MNoIndexUsedCnt

	// no validation rules for MNoIndexUsedSum

	// no validation rules for MNoGoodIndexUsedCnt

	// no validation rules for MNoGoodIndexUsedSum

	// no validation rules for MDocsReturnedCnt

	// no validation rules for MDocsReturnedSum

	// no validation rules for MDocsReturnedMin

	// no validation rules for MDocsReturnedMax

	// no validation rules for MDocsReturnedP99

	// no validation rules for MResponseLengthCnt

	// no validation rules for MResponseLengthSum

	// no validation rules for MResponseLengthMin

	// no validation rules for MResponseLengthMax

	// no validation rules for MResponseLengthP99

	// no validation rules for MDocsScannedCnt

	// no validation rules for MDocsScannedSum

	// no validation rules for MDocsScannedMin

	// no validation rules for MDocsScannedMax

	// no validation rules for MDocsScannedP99

	// no validation rules for MSharedBlksHitCnt

	// no validation rules for MSharedBlksHitSum

	// no validation rules for MSharedBlksReadCnt

	// no validation rules for MSharedBlksReadSum

	// no validation rules for MSharedBlksDirtiedCnt

	// no validation rules for MSharedBlksDirtiedSum

	// no validation rules for MSharedBlksWrittenCnt

	// no validation rules for MSharedBlksWrittenSum

	// no validation rules for MLocalBlksHitCnt

	// no validation rules for MLocalBlksHitSum

	// no validation rules for MLocalBlksReadCnt

	// no validation rules for MLocalBlksReadSum

	// no validation rules for MLocalBlksDirtiedCnt

	// no validation rules for MLocalBlksDirtiedSum

	// no validation rules for MLocalBlksWrittenCnt

	// no validation rules for MLocalBlksWrittenSum

	// no validation rules for MTempBlksReadCnt

	// no validation rules for MTempBlksReadSum

	// no validation rules for MTempBlksWrittenCnt

	// no validation rules for MTempBlksWrittenSum

	// no validation rules for MBlkReadTimeCnt

	// no validation rules for MBlkReadTimeSum

	// no validation rules for MBlkWriteTimeCnt

	// no validation rules for MBlkWriteTimeSum

	// no validation rules for MCpuUserTimeCnt

	// no validation rules for MCpuUserTimeSum

	// no validation rules for MCpuSysTimeCnt

	// no validation rules for MCpuSysTimeSum

	// no validation rules for CmdType

	// no validation rules for MPlansCallsSum

	// no validation rules for MPlansCallsCnt

	// no validation rules for MWalRecordsSum

	// no validation rules for MWalRecordsCnt

	// no validation rules for MWalFpiSum

	// no validation rules for MWalFpiCnt

	// no validation rules for MWalBytesSum

	// no validation rules for MWalBytesCnt

	// no validation rules for MPlanTimeSum

	// no validation rules for MPlanTimeCnt

	// no validation rules for MPlanTimeMin

	// no validation rules for MPlanTimeMax

	// no validation rules for TopQueryid

	// no validation rules for TopQuery

	// no validation rules for ApplicationName

	// no validation rules for Planid

	// no validation rules for QueryPlan

	if len(errors) > 0 {
		return MetricsBucketMultiError(errors)
	}

	return nil
}

// MetricsBucketMultiError is an error wrapping multiple validation errors
// returned by MetricsBucket.ValidateAll() if the designated constraints
// aren't met.
type MetricsBucketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsBucketMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsBucketMultiError) AllErrors() []error { return m }

// MetricsBucketValidationError is the validation error returned by
// MetricsBucket.Validate if the designated constraints aren't met.
type MetricsBucketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsBucketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsBucketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsBucketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsBucketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsBucketValidationError) ErrorName() string { return "MetricsBucketValidationError" }

// Error satisfies the builtin error interface
func (e MetricsBucketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsBucket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsBucketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsBucketValidationError{}

// Validate checks the field values on CollectResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CollectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CollectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CollectResponseMultiError, or nil if none found.
func (m *CollectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CollectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CollectResponseMultiError(errors)
	}

	return nil
}

// CollectResponseMultiError is an error wrapping multiple validation errors
// returned by CollectResponse.ValidateAll() if the designated constraints
// aren't met.
type CollectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CollectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CollectResponseMultiError) AllErrors() []error { return m }

// CollectResponseValidationError is the validation error returned by
// CollectResponse.Validate if the designated constraints aren't met.
type CollectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectResponseValidationError) ErrorName() string { return "CollectResponseValidationError" }

// Error satisfies the builtin error interface
func (e CollectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectResponseValidationError{}
