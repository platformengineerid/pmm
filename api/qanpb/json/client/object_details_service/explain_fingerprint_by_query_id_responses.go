// Code generated by go-swagger; DO NOT EDIT.

package object_details_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExplainFingerprintByQueryIDReader is a Reader for the ExplainFingerprintByQueryID structure.
type ExplainFingerprintByQueryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExplainFingerprintByQueryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExplainFingerprintByQueryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExplainFingerprintByQueryIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExplainFingerprintByQueryIDOK creates a ExplainFingerprintByQueryIDOK with default headers values
func NewExplainFingerprintByQueryIDOK() *ExplainFingerprintByQueryIDOK {
	return &ExplainFingerprintByQueryIDOK{}
}

/*
ExplainFingerprintByQueryIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExplainFingerprintByQueryIDOK struct {
	Payload *ExplainFingerprintByQueryIDOKBody
}

func (o *ExplainFingerprintByQueryIDOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/ExplainFingerprintByQueryID][%d] explainFingerprintByQueryIdOk  %+v", 200, o.Payload)
}

func (o *ExplainFingerprintByQueryIDOK) GetPayload() *ExplainFingerprintByQueryIDOKBody {
	return o.Payload
}

func (o *ExplainFingerprintByQueryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ExplainFingerprintByQueryIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExplainFingerprintByQueryIDDefault creates a ExplainFingerprintByQueryIDDefault with default headers values
func NewExplainFingerprintByQueryIDDefault(code int) *ExplainFingerprintByQueryIDDefault {
	return &ExplainFingerprintByQueryIDDefault{
		_statusCode: code,
	}
}

/*
ExplainFingerprintByQueryIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExplainFingerprintByQueryIDDefault struct {
	_statusCode int

	Payload *ExplainFingerprintByQueryIDDefaultBody
}

// Code gets the status code for the explain fingerprint by query ID default response
func (o *ExplainFingerprintByQueryIDDefault) Code() int {
	return o._statusCode
}

func (o *ExplainFingerprintByQueryIDDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/ExplainFingerprintByQueryID][%d] ExplainFingerprintByQueryID default  %+v", o._statusCode, o.Payload)
}

func (o *ExplainFingerprintByQueryIDDefault) GetPayload() *ExplainFingerprintByQueryIDDefaultBody {
	return o.Payload
}

func (o *ExplainFingerprintByQueryIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ExplainFingerprintByQueryIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExplainFingerprintByQueryIDBody ExplainFingerprintByQueryIDRequest get explain fingerprint for given query ID.
swagger:model ExplainFingerprintByQueryIDBody
*/
type ExplainFingerprintByQueryIDBody struct {
	// serviceid
	Serviceid string `json:"serviceid,omitempty"`

	// query id
	QueryID string `json:"query_id,omitempty"`
}

// Validate validates this explain fingerprint by query ID body
func (o *ExplainFingerprintByQueryIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this explain fingerprint by query ID body based on context it is used
func (o *ExplainFingerprintByQueryIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDBody) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExplainFingerprintByQueryIDDefaultBody explain fingerprint by query ID default body
swagger:model ExplainFingerprintByQueryIDDefaultBody
*/
type ExplainFingerprintByQueryIDDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this explain fingerprint by query ID default body
func (o *ExplainFingerprintByQueryIDDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExplainFingerprintByQueryIDDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this explain fingerprint by query ID default body based on the context it is used
func (o *ExplainFingerprintByQueryIDDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExplainFingerprintByQueryIDDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExplainFingerprintByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBody) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 explain fingerprint by query ID default body details items0
swagger:model ExplainFingerprintByQueryIDDefaultBodyDetailsItems0
*/
type ExplainFingerprintByQueryIDDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this explain fingerprint by query ID default body details items0
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this explain fingerprint by query ID default body details items0 based on context it is used
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExplainFingerprintByQueryIDOKBody ExplainFingerprintByQueryIDResponse is explain fingerprint and placeholders count for given query ID.
swagger:model ExplainFingerprintByQueryIDOKBody
*/
type ExplainFingerprintByQueryIDOKBody struct {
	// explain fingerprint
	ExplainFingerprint string `json:"explain_fingerprint,omitempty"`

	// placeholders count
	PlaceholdersCount int64 `json:"placeholders_count,omitempty"`
}

// Validate validates this explain fingerprint by query ID OK body
func (o *ExplainFingerprintByQueryIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this explain fingerprint by query ID OK body based on context it is used
func (o *ExplainFingerprintByQueryIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExplainFingerprintByQueryIDOKBody) UnmarshalBinary(b []byte) error {
	var res ExplainFingerprintByQueryIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
