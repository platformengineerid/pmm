// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/service/service.proto

package servicev1beta1

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

	_ = inventorypbv1.ServiceType(0)
)

// Validate checks the field values on UniversalService with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UniversalService) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UniversalService with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UniversalServiceMultiError, or nil if none found.
func (m *UniversalService) ValidateAll() error {
	return m.validate(true)
}

func (m *UniversalService) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceId

	// no validation rules for ServiceType

	// no validation rules for ServiceName

	// no validation rules for DatabaseName

	// no validation rules for NodeId

	// no validation rules for NodeName

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	// no validation rules for CustomLabels

	// no validation rules for ExternalGroup

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for Socket

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UniversalServiceValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UniversalServiceValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UniversalServiceValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UniversalServiceValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UniversalServiceValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UniversalServiceValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAgents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UniversalServiceValidationError{
						field:  fmt.Sprintf("Agents[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UniversalServiceValidationError{
						field:  fmt.Sprintf("Agents[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UniversalServiceValidationError{
					field:  fmt.Sprintf("Agents[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Status

	// no validation rules for Version

	if len(errors) > 0 {
		return UniversalServiceMultiError(errors)
	}

	return nil
}

// UniversalServiceMultiError is an error wrapping multiple validation errors
// returned by UniversalService.ValidateAll() if the designated constraints
// aren't met.
type UniversalServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UniversalServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UniversalServiceMultiError) AllErrors() []error { return m }

// UniversalServiceValidationError is the validation error returned by
// UniversalService.Validate if the designated constraints aren't met.
type UniversalServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UniversalServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UniversalServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UniversalServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UniversalServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UniversalServiceValidationError) ErrorName() string { return "UniversalServiceValidationError" }

// Error satisfies the builtin error interface
func (e UniversalServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUniversalService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UniversalServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UniversalServiceValidationError{}

// Validate checks the field values on ListServicesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListServicesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListServicesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListServicesRequestMultiError, or nil if none found.
func (m *ListServicesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListServicesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeId

	// no validation rules for ServiceType

	// no validation rules for ExternalGroup

	if len(errors) > 0 {
		return ListServicesRequestMultiError(errors)
	}

	return nil
}

// ListServicesRequestMultiError is an error wrapping multiple validation
// errors returned by ListServicesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListServicesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListServicesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListServicesRequestMultiError) AllErrors() []error { return m }

// ListServicesRequestValidationError is the validation error returned by
// ListServicesRequest.Validate if the designated constraints aren't met.
type ListServicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListServicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListServicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListServicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListServicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListServicesRequestValidationError) ErrorName() string {
	return "ListServicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListServicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListServicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListServicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListServicesRequestValidationError{}

// Validate checks the field values on ListServicesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListServicesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListServicesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListServicesResponseMultiError, or nil if none found.
func (m *ListServicesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListServicesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListServicesResponseValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListServicesResponseValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListServicesResponseValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListServicesResponseMultiError(errors)
	}

	return nil
}

// ListServicesResponseMultiError is an error wrapping multiple validation
// errors returned by ListServicesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListServicesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListServicesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListServicesResponseMultiError) AllErrors() []error { return m }

// ListServicesResponseValidationError is the validation error returned by
// ListServicesResponse.Validate if the designated constraints aren't met.
type ListServicesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListServicesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListServicesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListServicesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListServicesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListServicesResponseValidationError) ErrorName() string {
	return "ListServicesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListServicesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListServicesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListServicesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListServicesResponseValidationError{}
