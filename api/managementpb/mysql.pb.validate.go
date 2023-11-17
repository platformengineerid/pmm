// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/mysql.proto

package managementpbv1

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

	inventorypbv1 "github.com/percona/pmm/api/inventorypb/v1"
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

	_ = inventorypbv1.LogLevel(0)
)

// Validate checks the field values on AddMySQLRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddMySQLRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddMySQLRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddMySQLRequestMultiError, or nil if none found.
func (m *AddMySQLRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddMySQLRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeId

	// no validation rules for NodeName

	if all {
		switch v := interface{}(m.GetAddNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddMySQLRequestValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddMySQLRequestValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddMySQLRequestValidationError{
				field:  "AddNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		err := AddMySQLRequestValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for Socket

	if utf8.RuneCountInString(m.GetPmmAgentId()) < 1 {
		err := AddMySQLRequestValidationError{
			field:  "PmmAgentId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := AddMySQLRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Password

	// no validation rules for QanMysqlPerfschema

	// no validation rules for QanMysqlSlowlog

	// no validation rules for CustomLabels

	// no validation rules for SkipConnectionCheck

	// no validation rules for DisableCommentsParsing

	// no validation rules for MaxQueryLength

	// no validation rules for DisableQueryExamples

	// no validation rules for MaxSlowlogFileSize

	// no validation rules for Tls

	// no validation rules for TlsSkipVerify

	// no validation rules for TlsCa

	// no validation rules for TlsCert

	// no validation rules for TlsKey

	// no validation rules for TablestatsGroupTableLimit

	// no validation rules for MetricsMode

	// no validation rules for AgentPassword

	// no validation rules for LogLevel

	if len(errors) > 0 {
		return AddMySQLRequestMultiError(errors)
	}

	return nil
}

// AddMySQLRequestMultiError is an error wrapping multiple validation errors
// returned by AddMySQLRequest.ValidateAll() if the designated constraints
// aren't met.
type AddMySQLRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddMySQLRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddMySQLRequestMultiError) AllErrors() []error { return m }

// AddMySQLRequestValidationError is the validation error returned by
// AddMySQLRequest.Validate if the designated constraints aren't met.
type AddMySQLRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddMySQLRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddMySQLRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddMySQLRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddMySQLRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddMySQLRequestValidationError) ErrorName() string { return "AddMySQLRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddMySQLRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddMySQLRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddMySQLRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddMySQLRequestValidationError{}

// Validate checks the field values on AddMySQLResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddMySQLResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddMySQLResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddMySQLResponseMultiError, or nil if none found.
func (m *AddMySQLResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddMySQLResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddMySQLResponseValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMysqldExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "MysqldExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "MysqldExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMysqldExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddMySQLResponseValidationError{
				field:  "MysqldExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetQanMysqlPerfschema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "QanMysqlPerfschema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "QanMysqlPerfschema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQanMysqlPerfschema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddMySQLResponseValidationError{
				field:  "QanMysqlPerfschema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetQanMysqlSlowlog()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "QanMysqlSlowlog",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddMySQLResponseValidationError{
					field:  "QanMysqlSlowlog",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQanMysqlSlowlog()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddMySQLResponseValidationError{
				field:  "QanMysqlSlowlog",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TableCount

	if len(errors) > 0 {
		return AddMySQLResponseMultiError(errors)
	}

	return nil
}

// AddMySQLResponseMultiError is an error wrapping multiple validation errors
// returned by AddMySQLResponse.ValidateAll() if the designated constraints
// aren't met.
type AddMySQLResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddMySQLResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddMySQLResponseMultiError) AllErrors() []error { return m }

// AddMySQLResponseValidationError is the validation error returned by
// AddMySQLResponse.Validate if the designated constraints aren't met.
type AddMySQLResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddMySQLResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddMySQLResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddMySQLResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddMySQLResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddMySQLResponseValidationError) ErrorName() string { return "AddMySQLResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddMySQLResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddMySQLResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddMySQLResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddMySQLResponseValidationError{}
