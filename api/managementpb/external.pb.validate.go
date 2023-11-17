// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/external.proto

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
)

// Validate checks the field values on AddExternalRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddExternalRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddExternalRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddExternalRequestMultiError, or nil if none found.
func (m *AddExternalRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddExternalRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RunsOnNodeId

	// no validation rules for NodeName

	if all {
		switch v := interface{}(m.GetAddNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddExternalRequestValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddExternalRequestValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddExternalRequestValidationError{
				field:  "AddNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Address

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		err := AddExternalRequestValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Scheme

	// no validation rules for MetricsPath

	if val := m.GetListenPort(); val <= 0 || val >= 65536 {
		err := AddExternalRequestValidationError{
			field:  "ListenPort",
			reason: "value must be inside range (0, 65536)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NodeId

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	// no validation rules for CustomLabels

	// no validation rules for Group

	// no validation rules for MetricsMode

	// no validation rules for SkipConnectionCheck

	if len(errors) > 0 {
		return AddExternalRequestMultiError(errors)
	}

	return nil
}

// AddExternalRequestMultiError is an error wrapping multiple validation errors
// returned by AddExternalRequest.ValidateAll() if the designated constraints
// aren't met.
type AddExternalRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddExternalRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddExternalRequestMultiError) AllErrors() []error { return m }

// AddExternalRequestValidationError is the validation error returned by
// AddExternalRequest.Validate if the designated constraints aren't met.
type AddExternalRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddExternalRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddExternalRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddExternalRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddExternalRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddExternalRequestValidationError) ErrorName() string {
	return "AddExternalRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddExternalRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddExternalRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddExternalRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddExternalRequestValidationError{}

// Validate checks the field values on AddExternalResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddExternalResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddExternalResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddExternalResponseMultiError, or nil if none found.
func (m *AddExternalResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddExternalResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddExternalResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddExternalResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddExternalResponseValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExternalExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddExternalResponseValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddExternalResponseValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExternalExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddExternalResponseValidationError{
				field:  "ExternalExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddExternalResponseMultiError(errors)
	}

	return nil
}

// AddExternalResponseMultiError is an error wrapping multiple validation
// errors returned by AddExternalResponse.ValidateAll() if the designated
// constraints aren't met.
type AddExternalResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddExternalResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddExternalResponseMultiError) AllErrors() []error { return m }

// AddExternalResponseValidationError is the validation error returned by
// AddExternalResponse.Validate if the designated constraints aren't met.
type AddExternalResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddExternalResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddExternalResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddExternalResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddExternalResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddExternalResponseValidationError) ErrorName() string {
	return "AddExternalResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddExternalResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddExternalResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddExternalResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddExternalResponseValidationError{}
