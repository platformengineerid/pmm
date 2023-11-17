// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: agentlocalpb/agentlocal.proto

package agentlocalpbv1

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

// Validate checks the field values on ServerInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServerInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServerInfoMultiError, or
// nil if none found.
func (m *ServerInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	// no validation rules for InsecureTls

	// no validation rules for Connected

	// no validation rules for Version

	if all {
		switch v := interface{}(m.GetLatency()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerInfoValidationError{
					field:  "Latency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerInfoValidationError{
					field:  "Latency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLatency()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerInfoValidationError{
				field:  "Latency",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClockDrift()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerInfoValidationError{
					field:  "ClockDrift",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerInfoValidationError{
					field:  "ClockDrift",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClockDrift()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerInfoValidationError{
				field:  "ClockDrift",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServerInfoMultiError(errors)
	}

	return nil
}

// ServerInfoMultiError is an error wrapping multiple validation errors
// returned by ServerInfo.ValidateAll() if the designated constraints aren't met.
type ServerInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerInfoMultiError) AllErrors() []error { return m }

// ServerInfoValidationError is the validation error returned by
// ServerInfo.Validate if the designated constraints aren't met.
type ServerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoValidationError) ErrorName() string { return "ServerInfoValidationError" }

// Error satisfies the builtin error interface
func (e ServerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoValidationError{}

// Validate checks the field values on AgentInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AgentInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AgentInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AgentInfoMultiError, or nil
// if none found.
func (m *AgentInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *AgentInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AgentId

	// no validation rules for AgentType

	// no validation rules for Status

	// no validation rules for ListenPort

	// no validation rules for ProcessExecPath

	if len(errors) > 0 {
		return AgentInfoMultiError(errors)
	}

	return nil
}

// AgentInfoMultiError is an error wrapping multiple validation errors returned
// by AgentInfo.ValidateAll() if the designated constraints aren't met.
type AgentInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgentInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgentInfoMultiError) AllErrors() []error { return m }

// AgentInfoValidationError is the validation error returned by
// AgentInfo.Validate if the designated constraints aren't met.
type AgentInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgentInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgentInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgentInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgentInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgentInfoValidationError) ErrorName() string { return "AgentInfoValidationError" }

// Error satisfies the builtin error interface
func (e AgentInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgentInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgentInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgentInfoValidationError{}

// Validate checks the field values on StatusRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatusRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StatusRequestMultiError, or
// nil if none found.
func (m *StatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GetNetworkInfo

	if len(errors) > 0 {
		return StatusRequestMultiError(errors)
	}

	return nil
}

// StatusRequestMultiError is an error wrapping multiple validation errors
// returned by StatusRequest.ValidateAll() if the designated constraints
// aren't met.
type StatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatusRequestMultiError) AllErrors() []error { return m }

// StatusRequestValidationError is the validation error returned by
// StatusRequest.Validate if the designated constraints aren't met.
type StatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusRequestValidationError) ErrorName() string { return "StatusRequestValidationError" }

// Error satisfies the builtin error interface
func (e StatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusRequestValidationError{}

// Validate checks the field values on StatusResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatusResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StatusResponseMultiError,
// or nil if none found.
func (m *StatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AgentId

	// no validation rules for RunsOnNodeId

	if all {
		switch v := interface{}(m.GetServerInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatusResponseValidationError{
					field:  "ServerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatusResponseValidationError{
					field:  "ServerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServerInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusResponseValidationError{
				field:  "ServerInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAgentsInfo() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StatusResponseValidationError{
						field:  fmt.Sprintf("AgentsInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StatusResponseValidationError{
						field:  fmt.Sprintf("AgentsInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatusResponseValidationError{
					field:  fmt.Sprintf("AgentsInfo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ConfigFilepath

	// no validation rules for AgentVersion

	// no validation rules for NodeName

	// no validation rules for ConnectionUptime

	if len(errors) > 0 {
		return StatusResponseMultiError(errors)
	}

	return nil
}

// StatusResponseMultiError is an error wrapping multiple validation errors
// returned by StatusResponse.ValidateAll() if the designated constraints
// aren't met.
type StatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatusResponseMultiError) AllErrors() []error { return m }

// StatusResponseValidationError is the validation error returned by
// StatusResponse.Validate if the designated constraints aren't met.
type StatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusResponseValidationError) ErrorName() string { return "StatusResponseValidationError" }

// Error satisfies the builtin error interface
func (e StatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusResponseValidationError{}

// Validate checks the field values on ReloadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReloadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReloadRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReloadRequestMultiError, or
// nil if none found.
func (m *ReloadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReloadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReloadRequestMultiError(errors)
	}

	return nil
}

// ReloadRequestMultiError is an error wrapping multiple validation errors
// returned by ReloadRequest.ValidateAll() if the designated constraints
// aren't met.
type ReloadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReloadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReloadRequestMultiError) AllErrors() []error { return m }

// ReloadRequestValidationError is the validation error returned by
// ReloadRequest.Validate if the designated constraints aren't met.
type ReloadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReloadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReloadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReloadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReloadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReloadRequestValidationError) ErrorName() string { return "ReloadRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReloadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReloadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReloadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReloadRequestValidationError{}

// Validate checks the field values on ReloadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReloadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReloadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReloadResponseMultiError,
// or nil if none found.
func (m *ReloadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReloadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReloadResponseMultiError(errors)
	}

	return nil
}

// ReloadResponseMultiError is an error wrapping multiple validation errors
// returned by ReloadResponse.ValidateAll() if the designated constraints
// aren't met.
type ReloadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReloadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReloadResponseMultiError) AllErrors() []error { return m }

// ReloadResponseValidationError is the validation error returned by
// ReloadResponse.Validate if the designated constraints aren't met.
type ReloadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReloadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReloadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReloadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReloadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReloadResponseValidationError) ErrorName() string { return "ReloadResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReloadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReloadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReloadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReloadResponseValidationError{}
