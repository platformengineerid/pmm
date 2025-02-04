// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// RemoveChannelReader is a Reader for the RemoveChannel structure.
type RemoveChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveChannelOK creates a RemoveChannelOK with default headers values
func NewRemoveChannelOK() *RemoveChannelOK {
	return &RemoveChannelOK{}
}

/*
RemoveChannelOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveChannelOK struct {
	Payload interface{}
}

func (o *RemoveChannelOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/Remove][%d] removeChannelOk  %+v", 200, o.Payload)
}

func (o *RemoveChannelOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveChannelDefault creates a RemoveChannelDefault with default headers values
func NewRemoveChannelDefault(code int) *RemoveChannelDefault {
	return &RemoveChannelDefault{
		_statusCode: code,
	}
}

/*
RemoveChannelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveChannelDefault struct {
	_statusCode int

	Payload *RemoveChannelDefaultBody
}

// Code gets the status code for the remove channel default response
func (o *RemoveChannelDefault) Code() int {
	return o._statusCode
}

func (o *RemoveChannelDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/Remove][%d] RemoveChannel default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveChannelDefault) GetPayload() *RemoveChannelDefaultBody {
	return o.Payload
}

func (o *RemoveChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RemoveChannelDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RemoveChannelBody remove channel body
swagger:model RemoveChannelBody
*/
type RemoveChannelBody struct {
	// channel id
	ChannelID string `json:"channel_id,omitempty"`
}

// Validate validates this remove channel body
func (o *RemoveChannelBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove channel body based on context it is used
func (o *RemoveChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveChannelBody) UnmarshalBinary(b []byte) error {
	var res RemoveChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RemoveChannelDefaultBody remove channel default body
swagger:model RemoveChannelDefaultBody
*/
type RemoveChannelDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RemoveChannelDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this remove channel default body
func (o *RemoveChannelDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveChannelDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("RemoveChannel default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveChannel default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove channel default body based on the context it is used
func (o *RemoveChannelDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveChannelDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveChannel default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveChannel default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveChannelDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveChannelDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveChannelDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RemoveChannelDefaultBodyDetailsItems0 remove channel default body details items0
swagger:model RemoveChannelDefaultBodyDetailsItems0
*/
type RemoveChannelDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this remove channel default body details items0
func (o *RemoveChannelDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove channel default body details items0 based on context it is used
func (o *RemoveChannelDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveChannelDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveChannelDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RemoveChannelDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
