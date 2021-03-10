// Code generated by go-swagger; DO NOT EDIT.

package rds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

// DiscoverRDSReader is a Reader for the DiscoverRDS structure.
type DiscoverRDSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverRDSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoverRDSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDiscoverRDSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiscoverRDSOK creates a DiscoverRDSOK with default headers values
func NewDiscoverRDSOK() *DiscoverRDSOK {
	return &DiscoverRDSOK{}
}

/*DiscoverRDSOK handles this case with default header values.

A successful response.
*/
type DiscoverRDSOK struct {
	Payload *DiscoverRDSOKBody
}

func (o *DiscoverRDSOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/RDS/Discover][%d] discoverRdsOk  %+v", 200, o.Payload)
}

func (o *DiscoverRDSOK) GetPayload() *DiscoverRDSOKBody {
	return o.Payload
}

func (o *DiscoverRDSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DiscoverRDSOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverRDSDefault creates a DiscoverRDSDefault with default headers values
func NewDiscoverRDSDefault(code int) *DiscoverRDSDefault {
	return &DiscoverRDSDefault{
		_statusCode: code,
	}
}

/*DiscoverRDSDefault handles this case with default header values.

An unexpected error response.
*/
type DiscoverRDSDefault struct {
	_statusCode int

	Payload *DiscoverRDSDefaultBody
}

// Code gets the status code for the discover RDS default response
func (o *DiscoverRDSDefault) Code() int {
	return o._statusCode
}

func (o *DiscoverRDSDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/RDS/Discover][%d] DiscoverRDS default  %+v", o._statusCode, o.Payload)
}

func (o *DiscoverRDSDefault) GetPayload() *DiscoverRDSDefaultBody {
	return o.Payload
}

func (o *DiscoverRDSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DiscoverRDSDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DiscoverRDSBody discover RDS body
swagger:model DiscoverRDSBody
*/
type DiscoverRDSBody struct {

	// AWS Access key. Optional.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// AWS Secret key. Optional.
	AWSSecretKey string `json:"aws_secret_key,omitempty"`
}

// Validate validates this discover RDS body
func (o *DiscoverRDSBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverRDSBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverRDSBody) UnmarshalBinary(b []byte) error {
	var res DiscoverRDSBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DiscoverRDSDefaultBody discover RDS default body
swagger:model DiscoverRDSDefaultBody
*/
type DiscoverRDSDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this discover RDS default body
func (o *DiscoverRDSDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiscoverRDSDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("DiscoverRDS default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverRDSDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverRDSDefaultBody) UnmarshalBinary(b []byte) error {
	var res DiscoverRDSDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DiscoverRDSOKBody discover RDS OK body
swagger:model DiscoverRDSOKBody
*/
type DiscoverRDSOKBody struct {

	// rds instances
	RDSInstances []*RDSInstancesItems0 `json:"rds_instances"`
}

// Validate validates this discover RDS OK body
func (o *DiscoverRDSOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRDSInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiscoverRDSOKBody) validateRDSInstances(formats strfmt.Registry) error {

	if swag.IsZero(o.RDSInstances) { // not required
		return nil
	}

	for i := 0; i < len(o.RDSInstances); i++ {
		if swag.IsZero(o.RDSInstances[i]) { // not required
			continue
		}

		if o.RDSInstances[i] != nil {
			if err := o.RDSInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discoverRdsOk" + "." + "rds_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverRDSOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverRDSOKBody) UnmarshalBinary(b []byte) error {
	var res DiscoverRDSOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RDSInstancesItems0 DiscoverRDSInstance models an unique RDS instance for the list of instances returned by Discovery.
swagger:model RDSInstancesItems0
*/
type RDSInstancesItems0 struct {

	// AWS region.
	Region string `json:"region,omitempty"`

	// AWS availability zone.
	Az string `json:"az,omitempty"`

	// AWS instance ID.
	InstanceID string `json:"instance_id,omitempty"`

	// AWS instance class.
	NodeModel string `json:"node_model,omitempty"`

	// Address used to connect to it.
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// DiscoverRDSEngine describes supported RDS instance engines.
	// Enum: [DISCOVER_RDS_ENGINE_INVALID DISCOVER_RDS_MYSQL DISCOVER_RDS_POSTGRESQL]
	Engine *string `json:"engine,omitempty"`

	// Engine version.
	EngineVersion string `json:"engine_version,omitempty"`
}

// Validate validates this RDS instances items0
func (o *RDSInstancesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rdsInstancesItems0TypeEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISCOVER_RDS_ENGINE_INVALID","DISCOVER_RDS_MYSQL","DISCOVER_RDS_POSTGRESQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rdsInstancesItems0TypeEnginePropEnum = append(rdsInstancesItems0TypeEnginePropEnum, v)
	}
}

const (

	// RDSInstancesItems0EngineDISCOVERRDSENGINEINVALID captures enum value "DISCOVER_RDS_ENGINE_INVALID"
	RDSInstancesItems0EngineDISCOVERRDSENGINEINVALID string = "DISCOVER_RDS_ENGINE_INVALID"

	// RDSInstancesItems0EngineDISCOVERRDSMYSQL captures enum value "DISCOVER_RDS_MYSQL"
	RDSInstancesItems0EngineDISCOVERRDSMYSQL string = "DISCOVER_RDS_MYSQL"

	// RDSInstancesItems0EngineDISCOVERRDSPOSTGRESQL captures enum value "DISCOVER_RDS_POSTGRESQL"
	RDSInstancesItems0EngineDISCOVERRDSPOSTGRESQL string = "DISCOVER_RDS_POSTGRESQL"
)

// prop value enum
func (o *RDSInstancesItems0) validateEngineEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rdsInstancesItems0TypeEnginePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RDSInstancesItems0) validateEngine(formats strfmt.Registry) error {

	if swag.IsZero(o.Engine) { // not required
		return nil
	}

	// value enum
	if err := o.validateEngineEnum("engine", "body", *o.Engine); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RDSInstancesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RDSInstancesItems0) UnmarshalBinary(b []byte) error {
	var res RDSInstancesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
