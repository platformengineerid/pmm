// Code generated by go-swagger; DO NOT EDIT.

package services

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

// AddMongoDBServiceReader is a Reader for the AddMongoDBService structure.
type AddMongoDBServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMongoDBServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMongoDBServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMongoDBServiceOK creates a AddMongoDBServiceOK with default headers values
func NewAddMongoDBServiceOK() *AddMongoDBServiceOK {
	return &AddMongoDBServiceOK{}
}

/*
AddMongoDBServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddMongoDBServiceOK struct {
	Payload *AddMongoDBServiceOKBody
}

func (o *AddMongoDBServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMongoDB][%d] addMongoDbServiceOk  %+v", 200, o.Payload)
}

func (o *AddMongoDBServiceOK) GetPayload() *AddMongoDBServiceOKBody {
	return o.Payload
}

func (o *AddMongoDBServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddMongoDBServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMongoDBServiceDefault creates a AddMongoDBServiceDefault with default headers values
func NewAddMongoDBServiceDefault(code int) *AddMongoDBServiceDefault {
	return &AddMongoDBServiceDefault{
		_statusCode: code,
	}
}

/*
AddMongoDBServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddMongoDBServiceDefault struct {
	_statusCode int

	Payload *AddMongoDBServiceDefaultBody
}

// Code gets the status code for the add mongo DB service default response
func (o *AddMongoDBServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddMongoDBServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMongoDB][%d] AddMongoDBService default  %+v", o._statusCode, o.Payload)
}

func (o *AddMongoDBServiceDefault) GetPayload() *AddMongoDBServiceDefaultBody {
	return o.Payload
}

func (o *AddMongoDBServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddMongoDBServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddMongoDBServiceBody add mongo DB service body
swagger:model AddMongoDBServiceBody
*/
type AddMongoDBServiceBody struct {
	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add mongo DB service body
func (o *AddMongoDBServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add mongo DB service body based on context it is used
func (o *AddMongoDBServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBServiceDefaultBody add mongo DB service default body
swagger:model AddMongoDBServiceDefaultBody
*/
type AddMongoDBServiceDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddMongoDBServiceDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add mongo DB service default body
func (o *AddMongoDBServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBServiceDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddMongoDBService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddMongoDBService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add mongo DB service default body based on the context it is used
func (o *AddMongoDBServiceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBServiceDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddMongoDBService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddMongoDBService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBServiceDefaultBodyDetailsItems0 add mongo DB service default body details items0
swagger:model AddMongoDBServiceDefaultBodyDetailsItems0
*/
type AddMongoDBServiceDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add mongo DB service default body details items0
func (o *AddMongoDBServiceDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add mongo DB service default body details items0 based on context it is used
func (o *AddMongoDBServiceDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBServiceOKBody add mongo DB service OK body
swagger:model AddMongoDBServiceOKBody
*/
type AddMongoDBServiceOKBody struct {
	// mongodb
	Mongodb *AddMongoDBServiceOKBodyMongodb `json:"mongodb,omitempty"`
}

// Validate validates this add mongo DB service OK body
func (o *AddMongoDBServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBServiceOKBody) validateMongodb(formats strfmt.Registry) error {
	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	if o.Mongodb != nil {
		if err := o.Mongodb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbServiceOk" + "." + "mongodb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addMongoDbServiceOk" + "." + "mongodb")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add mongo DB service OK body based on the context it is used
func (o *AddMongoDBServiceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMongodb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBServiceOKBody) contextValidateMongodb(ctx context.Context, formats strfmt.Registry) error {
	if o.Mongodb != nil {
		if err := o.Mongodb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbServiceOk" + "." + "mongodb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addMongoDbServiceOk" + "." + "mongodb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddMongoDBServiceOKBodyMongodb MongoDBService represents a generic MongoDB instance.
swagger:model AddMongoDBServiceOKBodyMongodb
*/
type AddMongoDBServiceOKBodyMongodb struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// MongoDB version.
	Version string `json:"version,omitempty"`
}

// Validate validates this add mongo DB service OK body mongodb
func (o *AddMongoDBServiceOKBodyMongodb) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add mongo DB service OK body mongodb based on context it is used
func (o *AddMongoDBServiceOKBodyMongodb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceOKBodyMongodb) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceOKBodyMongodb) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceOKBodyMongodb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
