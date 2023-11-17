// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: qanpb/filters.proto

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

// Validate checks the field values on GetFilteredMetricsNamesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetFilteredMetricsNamesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetFilteredMetricsNamesRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetFilteredMetricsNamesRequestMultiError, or nil if none found.
func (m *GetFilteredMetricsNamesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetFilteredMetricsNamesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPeriodStartFrom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetFilteredMetricsNamesRequestValidationError{
					field:  "PeriodStartFrom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetFilteredMetricsNamesRequestValidationError{
					field:  "PeriodStartFrom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPeriodStartFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetFilteredMetricsNamesRequestValidationError{
				field:  "PeriodStartFrom",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPeriodStartTo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetFilteredMetricsNamesRequestValidationError{
					field:  "PeriodStartTo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetFilteredMetricsNamesRequestValidationError{
					field:  "PeriodStartTo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPeriodStartTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetFilteredMetricsNamesRequestValidationError{
				field:  "PeriodStartTo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MainMetricName

	for idx, item := range m.GetLabels() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetFilteredMetricsNamesRequestValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetFilteredMetricsNamesRequestValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetFilteredMetricsNamesRequestValidationError{
					field:  fmt.Sprintf("Labels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetFilteredMetricsNamesRequestMultiError(errors)
	}

	return nil
}

// GetFilteredMetricsNamesRequestMultiError is an error wrapping multiple
// validation errors returned by GetFilteredMetricsNamesRequest.ValidateAll()
// if the designated constraints aren't met.
type GetFilteredMetricsNamesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetFilteredMetricsNamesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetFilteredMetricsNamesRequestMultiError) AllErrors() []error { return m }

// GetFilteredMetricsNamesRequestValidationError is the validation error
// returned by GetFilteredMetricsNamesRequest.Validate if the designated
// constraints aren't met.
type GetFilteredMetricsNamesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetFilteredMetricsNamesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetFilteredMetricsNamesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetFilteredMetricsNamesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetFilteredMetricsNamesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetFilteredMetricsNamesRequestValidationError) ErrorName() string {
	return "GetFilteredMetricsNamesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetFilteredMetricsNamesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetFilteredMetricsNamesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetFilteredMetricsNamesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetFilteredMetricsNamesRequestValidationError{}

// Validate checks the field values on GetFilteredMetricsNamesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetFilteredMetricsNamesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetFilteredMetricsNamesResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetFilteredMetricsNamesResponseMultiError, or nil if none found.
func (m *GetFilteredMetricsNamesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetFilteredMetricsNamesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetLabels()))
		i := 0
		for key := range m.GetLabels() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetLabels()[key]
			_ = val

			// no validation rules for Labels[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, GetFilteredMetricsNamesResponseValidationError{
							field:  fmt.Sprintf("Labels[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, GetFilteredMetricsNamesResponseValidationError{
							field:  fmt.Sprintf("Labels[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return GetFilteredMetricsNamesResponseValidationError{
						field:  fmt.Sprintf("Labels[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return GetFilteredMetricsNamesResponseMultiError(errors)
	}

	return nil
}

// GetFilteredMetricsNamesResponseMultiError is an error wrapping multiple
// validation errors returned by GetFilteredMetricsNamesResponse.ValidateAll()
// if the designated constraints aren't met.
type GetFilteredMetricsNamesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetFilteredMetricsNamesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetFilteredMetricsNamesResponseMultiError) AllErrors() []error { return m }

// GetFilteredMetricsNamesResponseValidationError is the validation error
// returned by GetFilteredMetricsNamesResponse.Validate if the designated
// constraints aren't met.
type GetFilteredMetricsNamesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetFilteredMetricsNamesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetFilteredMetricsNamesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetFilteredMetricsNamesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetFilteredMetricsNamesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetFilteredMetricsNamesResponseValidationError) ErrorName() string {
	return "GetFilteredMetricsNamesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetFilteredMetricsNamesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetFilteredMetricsNamesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetFilteredMetricsNamesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetFilteredMetricsNamesResponseValidationError{}

// Validate checks the field values on ListLabels with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListLabels) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListLabels with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListLabelsMultiError, or
// nil if none found.
func (m *ListLabels) ValidateAll() error {
	return m.validate(true)
}

func (m *ListLabels) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetName() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListLabelsValidationError{
						field:  fmt.Sprintf("Name[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListLabelsValidationError{
						field:  fmt.Sprintf("Name[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListLabelsValidationError{
					field:  fmt.Sprintf("Name[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListLabelsMultiError(errors)
	}

	return nil
}

// ListLabelsMultiError is an error wrapping multiple validation errors
// returned by ListLabels.ValidateAll() if the designated constraints aren't met.
type ListLabelsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListLabelsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListLabelsMultiError) AllErrors() []error { return m }

// ListLabelsValidationError is the validation error returned by
// ListLabels.Validate if the designated constraints aren't met.
type ListLabelsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLabelsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLabelsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLabelsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLabelsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLabelsValidationError) ErrorName() string { return "ListLabelsValidationError" }

// Error satisfies the builtin error interface
func (e ListLabelsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLabels.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLabelsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLabelsValidationError{}

// Validate checks the field values on Values with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Values) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Values with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ValuesMultiError, or nil if none found.
func (m *Values) ValidateAll() error {
	return m.validate(true)
}

func (m *Values) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	// no validation rules for MainMetricPercent

	// no validation rules for MainMetricPerSec

	if len(errors) > 0 {
		return ValuesMultiError(errors)
	}

	return nil
}

// ValuesMultiError is an error wrapping multiple validation errors returned by
// Values.ValidateAll() if the designated constraints aren't met.
type ValuesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValuesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValuesMultiError) AllErrors() []error { return m }

// ValuesValidationError is the validation error returned by Values.Validate if
// the designated constraints aren't met.
type ValuesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValuesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValuesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValuesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValuesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValuesValidationError) ErrorName() string { return "ValuesValidationError" }

// Error satisfies the builtin error interface
func (e ValuesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValues.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValuesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValuesValidationError{}
