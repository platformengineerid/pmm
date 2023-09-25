// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/dump/dump.proto

package dumpv1beta1

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

// Validate checks the field values on Dump with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Dump) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dump with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DumpMultiError, or nil if none found.
func (m *Dump) ValidateAll() error {
	return m.validate(true)
}

func (m *Dump) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DumpId

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetStartTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DumpValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DumpValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DumpValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DumpValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DumpValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DumpValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DumpValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DumpValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DumpValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DumpMultiError(errors)
	}

	return nil
}

// DumpMultiError is an error wrapping multiple validation errors returned by
// Dump.ValidateAll() if the designated constraints aren't met.
type DumpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DumpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DumpMultiError) AllErrors() []error { return m }

// DumpValidationError is the validation error returned by Dump.Validate if the
// designated constraints aren't met.
type DumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DumpValidationError) ErrorName() string { return "DumpValidationError" }

// Error satisfies the builtin error interface
func (e DumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DumpValidationError{}

// Validate checks the field values on StartDumpRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StartDumpRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartDumpRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StartDumpRequestMultiError, or nil if none found.
func (m *StartDumpRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StartDumpRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStartTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StartDumpRequestValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StartDumpRequestValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StartDumpRequestValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StartDumpRequestValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StartDumpRequestValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StartDumpRequestValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ExportQan

	// no validation rules for IgnoreLoad

	if len(errors) > 0 {
		return StartDumpRequestMultiError(errors)
	}

	return nil
}

// StartDumpRequestMultiError is an error wrapping multiple validation errors
// returned by StartDumpRequest.ValidateAll() if the designated constraints
// aren't met.
type StartDumpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartDumpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartDumpRequestMultiError) AllErrors() []error { return m }

// StartDumpRequestValidationError is the validation error returned by
// StartDumpRequest.Validate if the designated constraints aren't met.
type StartDumpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartDumpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartDumpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartDumpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartDumpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartDumpRequestValidationError) ErrorName() string { return "StartDumpRequestValidationError" }

// Error satisfies the builtin error interface
func (e StartDumpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartDumpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartDumpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartDumpRequestValidationError{}

// Validate checks the field values on StartDumpResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StartDumpResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartDumpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StartDumpResponseMultiError, or nil if none found.
func (m *StartDumpResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StartDumpResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DumpId

	if len(errors) > 0 {
		return StartDumpResponseMultiError(errors)
	}

	return nil
}

// StartDumpResponseMultiError is an error wrapping multiple validation errors
// returned by StartDumpResponse.ValidateAll() if the designated constraints
// aren't met.
type StartDumpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartDumpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartDumpResponseMultiError) AllErrors() []error { return m }

// StartDumpResponseValidationError is the validation error returned by
// StartDumpResponse.Validate if the designated constraints aren't met.
type StartDumpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartDumpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartDumpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartDumpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartDumpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartDumpResponseValidationError) ErrorName() string {
	return "StartDumpResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StartDumpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartDumpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartDumpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartDumpResponseValidationError{}

// Validate checks the field values on ListDumpsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListDumpsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDumpsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDumpsRequestMultiError, or nil if none found.
func (m *ListDumpsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDumpsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListDumpsRequestMultiError(errors)
	}

	return nil
}

// ListDumpsRequestMultiError is an error wrapping multiple validation errors
// returned by ListDumpsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListDumpsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDumpsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDumpsRequestMultiError) AllErrors() []error { return m }

// ListDumpsRequestValidationError is the validation error returned by
// ListDumpsRequest.Validate if the designated constraints aren't met.
type ListDumpsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDumpsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDumpsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDumpsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDumpsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDumpsRequestValidationError) ErrorName() string { return "ListDumpsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListDumpsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDumpsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDumpsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDumpsRequestValidationError{}

// Validate checks the field values on ListDumpsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListDumpsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDumpsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDumpsResponseMultiError, or nil if none found.
func (m *ListDumpsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDumpsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDumps() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListDumpsResponseValidationError{
						field:  fmt.Sprintf("Dumps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListDumpsResponseValidationError{
						field:  fmt.Sprintf("Dumps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDumpsResponseValidationError{
					field:  fmt.Sprintf("Dumps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListDumpsResponseMultiError(errors)
	}

	return nil
}

// ListDumpsResponseMultiError is an error wrapping multiple validation errors
// returned by ListDumpsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListDumpsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDumpsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDumpsResponseMultiError) AllErrors() []error { return m }

// ListDumpsResponseValidationError is the validation error returned by
// ListDumpsResponse.Validate if the designated constraints aren't met.
type ListDumpsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDumpsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDumpsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDumpsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDumpsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDumpsResponseValidationError) ErrorName() string {
	return "ListDumpsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListDumpsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDumpsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDumpsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDumpsResponseValidationError{}

// Validate checks the field values on DeleteDumpRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteDumpRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDumpRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDumpRequestMultiError, or nil if none found.
func (m *DeleteDumpRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDumpRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DumpId

	if len(errors) > 0 {
		return DeleteDumpRequestMultiError(errors)
	}

	return nil
}

// DeleteDumpRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteDumpRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteDumpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDumpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDumpRequestMultiError) AllErrors() []error { return m }

// DeleteDumpRequestValidationError is the validation error returned by
// DeleteDumpRequest.Validate if the designated constraints aren't met.
type DeleteDumpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDumpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDumpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDumpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDumpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDumpRequestValidationError) ErrorName() string {
	return "DeleteDumpRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDumpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDumpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDumpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDumpRequestValidationError{}

// Validate checks the field values on DeleteDumpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteDumpResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDumpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDumpResponseMultiError, or nil if none found.
func (m *DeleteDumpResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDumpResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteDumpResponseMultiError(errors)
	}

	return nil
}

// DeleteDumpResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteDumpResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteDumpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDumpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDumpResponseMultiError) AllErrors() []error { return m }

// DeleteDumpResponseValidationError is the validation error returned by
// DeleteDumpResponse.Validate if the designated constraints aren't met.
type DeleteDumpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDumpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDumpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDumpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDumpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDumpResponseValidationError) ErrorName() string {
	return "DeleteDumpResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDumpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDumpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDumpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDumpResponseValidationError{}

// Validate checks the field values on GetDumpLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDumpLogsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDumpLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDumpLogsRequestMultiError, or nil if none found.
func (m *GetDumpLogsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDumpLogsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DumpId

	// no validation rules for Offset

	// no validation rules for Limit

	// no validation rules for RestoreId

	if len(errors) > 0 {
		return GetDumpLogsRequestMultiError(errors)
	}

	return nil
}

// GetDumpLogsRequestMultiError is an error wrapping multiple validation errors
// returned by GetDumpLogsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetDumpLogsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDumpLogsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDumpLogsRequestMultiError) AllErrors() []error { return m }

// GetDumpLogsRequestValidationError is the validation error returned by
// GetDumpLogsRequest.Validate if the designated constraints aren't met.
type GetDumpLogsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDumpLogsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDumpLogsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDumpLogsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDumpLogsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDumpLogsRequestValidationError) ErrorName() string {
	return "GetDumpLogsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDumpLogsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDumpLogsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDumpLogsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDumpLogsRequestValidationError{}

// Validate checks the field values on GetDumpLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDumpLogsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDumpLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDumpLogsResponseMultiError, or nil if none found.
func (m *GetDumpLogsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDumpLogsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetDumpLogsResponseMultiError(errors)
	}

	return nil
}

// GetDumpLogsResponseMultiError is an error wrapping multiple validation
// errors returned by GetDumpLogsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetDumpLogsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDumpLogsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDumpLogsResponseMultiError) AllErrors() []error { return m }

// GetDumpLogsResponseValidationError is the validation error returned by
// GetDumpLogsResponse.Validate if the designated constraints aren't met.
type GetDumpLogsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDumpLogsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDumpLogsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDumpLogsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDumpLogsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDumpLogsResponseValidationError) ErrorName() string {
	return "GetDumpLogsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDumpLogsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDumpLogsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDumpLogsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDumpLogsResponseValidationError{}
