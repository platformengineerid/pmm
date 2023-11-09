// Code generated by go-swagger; DO NOT EDIT.

package azure_database

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

// DiscoverAzureDatabaseReader is a Reader for the DiscoverAzureDatabase structure.
type DiscoverAzureDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverAzureDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoverAzureDatabaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDiscoverAzureDatabaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiscoverAzureDatabaseOK creates a DiscoverAzureDatabaseOK with default headers values
func NewDiscoverAzureDatabaseOK() *DiscoverAzureDatabaseOK {
	return &DiscoverAzureDatabaseOK{}
}

/*
DiscoverAzureDatabaseOK describes a response with status code 200, with default header values.

A successful response.
*/
type DiscoverAzureDatabaseOK struct {
	Payload *DiscoverAzureDatabaseOKBody
}

func (o *DiscoverAzureDatabaseOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/azure/AzureDatabase/Discover][%d] discoverAzureDatabaseOk  %+v", 200, o.Payload)
}

func (o *DiscoverAzureDatabaseOK) GetPayload() *DiscoverAzureDatabaseOKBody {
	return o.Payload
}

func (o *DiscoverAzureDatabaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(DiscoverAzureDatabaseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverAzureDatabaseDefault creates a DiscoverAzureDatabaseDefault with default headers values
func NewDiscoverAzureDatabaseDefault(code int) *DiscoverAzureDatabaseDefault {
	return &DiscoverAzureDatabaseDefault{
		_statusCode: code,
	}
}

/*
DiscoverAzureDatabaseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DiscoverAzureDatabaseDefault struct {
	_statusCode int

	Payload *DiscoverAzureDatabaseDefaultBody
}

// Code gets the status code for the discover azure database default response
func (o *DiscoverAzureDatabaseDefault) Code() int {
	return o._statusCode
}

func (o *DiscoverAzureDatabaseDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/azure/AzureDatabase/Discover][%d] DiscoverAzureDatabase default  %+v", o._statusCode, o.Payload)
}

func (o *DiscoverAzureDatabaseDefault) GetPayload() *DiscoverAzureDatabaseDefaultBody {
	return o.Payload
}

func (o *DiscoverAzureDatabaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(DiscoverAzureDatabaseDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DiscoverAzureDatabaseBody DiscoverAzureDatabaseRequest discover azure databases request.
swagger:model DiscoverAzureDatabaseBody
*/
type DiscoverAzureDatabaseBody struct {
	// Azure client ID.
	AzureClientID string `json:"azure_client_id,omitempty"`

	// Azure client secret.
	AzureClientSecret string `json:"azure_client_secret,omitempty"`

	// Azure tanant ID.
	AzureTenantID string `json:"azure_tenant_id,omitempty"`

	// Azure subscription ID.
	AzureSubscriptionID string `json:"azure_subscription_id,omitempty"`
}

// Validate validates this discover azure database body
func (o *DiscoverAzureDatabaseBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this discover azure database body based on context it is used
func (o *DiscoverAzureDatabaseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverAzureDatabaseBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverAzureDatabaseBody) UnmarshalBinary(b []byte) error {
	var res DiscoverAzureDatabaseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DiscoverAzureDatabaseDefaultBody discover azure database default body
swagger:model DiscoverAzureDatabaseDefaultBody
*/
type DiscoverAzureDatabaseDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DiscoverAzureDatabaseDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this discover azure database default body
func (o *DiscoverAzureDatabaseDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiscoverAzureDatabaseDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("DiscoverAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DiscoverAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this discover azure database default body based on the context it is used
func (o *DiscoverAzureDatabaseDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiscoverAzureDatabaseDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DiscoverAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DiscoverAzureDatabase default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverAzureDatabaseDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverAzureDatabaseDefaultBody) UnmarshalBinary(b []byte) error {
	var res DiscoverAzureDatabaseDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DiscoverAzureDatabaseDefaultBodyDetailsItems0 discover azure database default body details items0
swagger:model DiscoverAzureDatabaseDefaultBodyDetailsItems0
*/
type DiscoverAzureDatabaseDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this discover azure database default body details items0
func (o *DiscoverAzureDatabaseDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this discover azure database default body details items0 based on context it is used
func (o *DiscoverAzureDatabaseDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverAzureDatabaseDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverAzureDatabaseDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DiscoverAzureDatabaseDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DiscoverAzureDatabaseOKBody DiscoverAzureDatabaseResponse discover azure databases response.
swagger:model DiscoverAzureDatabaseOKBody
*/
type DiscoverAzureDatabaseOKBody struct {
	// azure database instance
	AzureDatabaseInstance []*DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0 `json:"azure_database_instance"`
}

// Validate validates this discover azure database OK body
func (o *DiscoverAzureDatabaseOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAzureDatabaseInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiscoverAzureDatabaseOKBody) validateAzureDatabaseInstance(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureDatabaseInstance) { // not required
		return nil
	}

	for i := 0; i < len(o.AzureDatabaseInstance); i++ {
		if swag.IsZero(o.AzureDatabaseInstance[i]) { // not required
			continue
		}

		if o.AzureDatabaseInstance[i] != nil {
			if err := o.AzureDatabaseInstance[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discoverAzureDatabaseOk" + "." + "azure_database_instance" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discoverAzureDatabaseOk" + "." + "azure_database_instance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this discover azure database OK body based on the context it is used
func (o *DiscoverAzureDatabaseOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAzureDatabaseInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiscoverAzureDatabaseOKBody) contextValidateAzureDatabaseInstance(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.AzureDatabaseInstance); i++ {
		if o.AzureDatabaseInstance[i] != nil {
			if err := o.AzureDatabaseInstance[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discoverAzureDatabaseOk" + "." + "azure_database_instance" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discoverAzureDatabaseOk" + "." + "azure_database_instance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverAzureDatabaseOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverAzureDatabaseOKBody) UnmarshalBinary(b []byte) error {
	var res DiscoverAzureDatabaseOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0 DiscoverAzureDatabaseInstance models an unique Azure Database instance for the list of instances returned by Discovery.
swagger:model DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0
*/
type DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0 struct {
	// Azure database instance ID.
	InstanceID string `json:"instance_id,omitempty"`

	// Azure database location.
	Region string `json:"region,omitempty"`

	// Azure database server name.
	ServiceName string `json:"service_name,omitempty"`

	// Database username.
	Username string `json:"username,omitempty"`

	// Address used to connect to it.
	Address string `json:"address,omitempty"`

	// Azure Resource group.
	AzureResourceGroup string `json:"azure_resource_group,omitempty"`

	// Environment tag.
	Environment string `json:"environment,omitempty"`

	// DiscoverAzureDatabaseType describes supported Azure Database instance engines.
	//
	//  - DISCOVER_AZURE_DATABASE_TYPE_MYSQL: MySQL type: microsoft.dbformysql or MariaDB type: microsoft.dbformariadb
	//  - DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL: PostgreSQL type: microsoft.dbformysql
	// Enum: [DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED DISCOVER_AZURE_DATABASE_TYPE_MYSQL DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL]
	Type *string `json:"type,omitempty"`

	// Azure database availability zone.
	Az string `json:"az,omitempty"`

	// Represents a purchasable Stock Keeping Unit (SKU) under a product.
	// https://docs.microsoft.com/en-us/partner-center/develop/product-resources#sku.
	NodeModel string `json:"node_model,omitempty"`
}

// Validate validates this discover azure database OK body azure database instance items0
func (o *DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var discoverAzureDatabaseOkBodyAzureDatabaseInstanceItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED","DISCOVER_AZURE_DATABASE_TYPE_MYSQL","DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discoverAzureDatabaseOkBodyAzureDatabaseInstanceItems0TypeTypePropEnum = append(discoverAzureDatabaseOkBodyAzureDatabaseInstanceItems0TypeTypePropEnum, v)
	}
}

const (

	// DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0TypeDISCOVERAZUREDATABASETYPEUNSPECIFIED captures enum value "DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED"
	DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0TypeDISCOVERAZUREDATABASETYPEUNSPECIFIED string = "DISCOVER_AZURE_DATABASE_TYPE_UNSPECIFIED"

	// DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0TypeDISCOVERAZUREDATABASETYPEMYSQL captures enum value "DISCOVER_AZURE_DATABASE_TYPE_MYSQL"
	DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0TypeDISCOVERAZUREDATABASETYPEMYSQL string = "DISCOVER_AZURE_DATABASE_TYPE_MYSQL"

	// DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0TypeDISCOVERAZUREDATABASETYPEPOSTGRESQL captures enum value "DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"
	DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0TypeDISCOVERAZUREDATABASETYPEPOSTGRESQL string = "DISCOVER_AZURE_DATABASE_TYPE_POSTGRESQL"
)

// prop value enum
func (o *DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, discoverAzureDatabaseOkBodyAzureDatabaseInstanceItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this discover azure database OK body azure database instance items0 based on context it is used
func (o *DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0) UnmarshalBinary(b []byte) error {
	var res DiscoverAzureDatabaseOKBodyAzureDatabaseInstanceItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
