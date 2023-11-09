// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListActiveServiceTypesReader is a Reader for the ListActiveServiceTypes structure.
type ListActiveServiceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListActiveServiceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListActiveServiceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListActiveServiceTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListActiveServiceTypesOK creates a ListActiveServiceTypesOK with default headers values
func NewListActiveServiceTypesOK() *ListActiveServiceTypesOK {
	return &ListActiveServiceTypesOK{}
}

/*
ListActiveServiceTypesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListActiveServiceTypesOK struct {
	Payload *ListActiveServiceTypesOKBody
}

func (o *ListActiveServiceTypesOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/ListTypes][%d] listActiveServiceTypesOk  %+v", 200, o.Payload)
}

func (o *ListActiveServiceTypesOK) GetPayload() *ListActiveServiceTypesOKBody {
	return o.Payload
}

func (o *ListActiveServiceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListActiveServiceTypesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListActiveServiceTypesDefault creates a ListActiveServiceTypesDefault with default headers values
func NewListActiveServiceTypesDefault(code int) *ListActiveServiceTypesDefault {
	return &ListActiveServiceTypesDefault{
		_statusCode: code,
	}
}

/*
ListActiveServiceTypesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListActiveServiceTypesDefault struct {
	_statusCode int

	Payload *ListActiveServiceTypesDefaultBody
}

// Code gets the status code for the list active service types default response
func (o *ListActiveServiceTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListActiveServiceTypesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/ListTypes][%d] ListActiveServiceTypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListActiveServiceTypesDefault) GetPayload() *ListActiveServiceTypesDefaultBody {
	return o.Payload
}

func (o *ListActiveServiceTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListActiveServiceTypesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListActiveServiceTypesDefaultBody list active service types default body
swagger:model ListActiveServiceTypesDefaultBody
*/
type ListActiveServiceTypesDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListActiveServiceTypesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list active service types default body
func (o *ListActiveServiceTypesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListActiveServiceTypesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListActiveServiceTypes default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListActiveServiceTypes default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list active service types default body based on the context it is used
func (o *ListActiveServiceTypesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListActiveServiceTypesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListActiveServiceTypes default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListActiveServiceTypes default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListActiveServiceTypesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListActiveServiceTypesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListActiveServiceTypesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListActiveServiceTypesDefaultBodyDetailsItems0 list active service types default body details items0
swagger:model ListActiveServiceTypesDefaultBodyDetailsItems0
*/
type ListActiveServiceTypesDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list active service types default body details items0
func (o *ListActiveServiceTypesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list active service types default body details items0 based on context it is used
func (o *ListActiveServiceTypesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListActiveServiceTypesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListActiveServiceTypesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListActiveServiceTypesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListActiveServiceTypesOKBody list active service types OK body
swagger:model ListActiveServiceTypesOKBody
*/
type ListActiveServiceTypesOKBody struct {
	// service types
	ServiceTypes []*string `json:"service_types"`
}

// Validate validates this list active service types OK body
func (o *ListActiveServiceTypesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServiceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listActiveServiceTypesOkBodyServiceTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SERVICE_TYPE_UNSPECIFIED","MYSQL_SERVICE","MONGODB_SERVICE","POSTGRESQL_SERVICE","PROXYSQL_SERVICE","HAPROXY_SERVICE","EXTERNAL_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listActiveServiceTypesOkBodyServiceTypesItemsEnum = append(listActiveServiceTypesOkBodyServiceTypesItemsEnum, v)
	}
}

func (o *ListActiveServiceTypesOKBody) validateServiceTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listActiveServiceTypesOkBodyServiceTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListActiveServiceTypesOKBody) validateServiceTypes(formats strfmt.Registry) error {
	if swag.IsZero(o.ServiceTypes) { // not required
		return nil
	}

	for i := 0; i < len(o.ServiceTypes); i++ {
		if swag.IsZero(o.ServiceTypes[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateServiceTypesItemsEnum("listActiveServiceTypesOk"+"."+"service_types"+"."+strconv.Itoa(i), "body", *o.ServiceTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this list active service types OK body based on context it is used
func (o *ListActiveServiceTypesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListActiveServiceTypesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListActiveServiceTypesOKBody) UnmarshalBinary(b []byte) error {
	var res ListActiveServiceTypesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
