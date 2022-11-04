// Code generated by go-swagger; DO NOT EDIT.

package rules

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

// CreateAlertRuleReader is a Reader for the CreateAlertRule structure.
type CreateAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAlertRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAlertRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAlertRuleOK creates a CreateAlertRuleOK with default headers values
func NewCreateAlertRuleOK() *CreateAlertRuleOK {
	return &CreateAlertRuleOK{}
}

/*
CreateAlertRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAlertRuleOK struct {
	Payload *CreateAlertRuleOKBody
}

func (o *CreateAlertRuleOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Create][%d] createAlertRuleOk  %+v", 200, o.Payload)
}

func (o *CreateAlertRuleOK) GetPayload() *CreateAlertRuleOKBody {
	return o.Payload
}

func (o *CreateAlertRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(CreateAlertRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertRuleDefault creates a CreateAlertRuleDefault with default headers values
func NewCreateAlertRuleDefault(code int) *CreateAlertRuleDefault {
	return &CreateAlertRuleDefault{
		_statusCode: code,
	}
}

/*
CreateAlertRuleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAlertRuleDefault struct {
	_statusCode int

	Payload *CreateAlertRuleDefaultBody
}

// Code gets the status code for the create alert rule default response
func (o *CreateAlertRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateAlertRuleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Create][%d] CreateAlertRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAlertRuleDefault) GetPayload() *CreateAlertRuleDefaultBody {
	return o.Payload
}

func (o *CreateAlertRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(CreateAlertRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateAlertRuleBody create alert rule body
swagger:model CreateAlertRuleBody
*/
type CreateAlertRuleBody struct {
	// Template name. Can't be specified simultaneously with source_rule_id.
	TemplateName string `json:"template_name,omitempty"`

	// ID of the rule that will be used as source. Can't be specified simultaneously with template_name.
	SourceRuleID string `json:"source_rule_id,omitempty"`

	// Rule name.
	Name string `json:"name,omitempty"`

	// New rule status.
	Disabled bool `json:"disabled,omitempty"`

	// Rule parameters. All template parameters should be set.
	Params []*CreateAlertRuleParamsBodyParamsItems0 `json:"params"`

	// Rule duration. Should be set.
	For string `json:"for,omitempty"`

	// Severity represents severity level of the check result or alert.
	// Enum: [SEVERITY_INVALID SEVERITY_EMERGENCY SEVERITY_ALERT SEVERITY_CRITICAL SEVERITY_ERROR SEVERITY_WARNING SEVERITY_NOTICE SEVERITY_INFO SEVERITY_DEBUG]
	Severity *string `json:"severity,omitempty"`

	// All custom labels to add or remove (with empty values) to default labels from template.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Filters. Should be set.
	Filters []*CreateAlertRuleParamsBodyFiltersItems0 `json:"filters"`

	// Channels. Should be set.
	ChannelIds []string `json:"channel_ids"`
}

// Validate validates this create alert rule body
func (o *CreateAlertRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAlertRuleBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var createAlertRuleBodyTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEVERITY_INVALID","SEVERITY_EMERGENCY","SEVERITY_ALERT","SEVERITY_CRITICAL","SEVERITY_ERROR","SEVERITY_WARNING","SEVERITY_NOTICE","SEVERITY_INFO","SEVERITY_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAlertRuleBodyTypeSeverityPropEnum = append(createAlertRuleBodyTypeSeverityPropEnum, v)
	}
}

const (

	// CreateAlertRuleBodySeveritySEVERITYINVALID captures enum value "SEVERITY_INVALID"
	CreateAlertRuleBodySeveritySEVERITYINVALID string = "SEVERITY_INVALID"

	// CreateAlertRuleBodySeveritySEVERITYEMERGENCY captures enum value "SEVERITY_EMERGENCY"
	CreateAlertRuleBodySeveritySEVERITYEMERGENCY string = "SEVERITY_EMERGENCY"

	// CreateAlertRuleBodySeveritySEVERITYALERT captures enum value "SEVERITY_ALERT"
	CreateAlertRuleBodySeveritySEVERITYALERT string = "SEVERITY_ALERT"

	// CreateAlertRuleBodySeveritySEVERITYCRITICAL captures enum value "SEVERITY_CRITICAL"
	CreateAlertRuleBodySeveritySEVERITYCRITICAL string = "SEVERITY_CRITICAL"

	// CreateAlertRuleBodySeveritySEVERITYERROR captures enum value "SEVERITY_ERROR"
	CreateAlertRuleBodySeveritySEVERITYERROR string = "SEVERITY_ERROR"

	// CreateAlertRuleBodySeveritySEVERITYWARNING captures enum value "SEVERITY_WARNING"
	CreateAlertRuleBodySeveritySEVERITYWARNING string = "SEVERITY_WARNING"

	// CreateAlertRuleBodySeveritySEVERITYNOTICE captures enum value "SEVERITY_NOTICE"
	CreateAlertRuleBodySeveritySEVERITYNOTICE string = "SEVERITY_NOTICE"

	// CreateAlertRuleBodySeveritySEVERITYINFO captures enum value "SEVERITY_INFO"
	CreateAlertRuleBodySeveritySEVERITYINFO string = "SEVERITY_INFO"

	// CreateAlertRuleBodySeveritySEVERITYDEBUG captures enum value "SEVERITY_DEBUG"
	CreateAlertRuleBodySeveritySEVERITYDEBUG string = "SEVERITY_DEBUG"
)

// prop value enum
func (o *CreateAlertRuleBody) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAlertRuleBodyTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateAlertRuleBody) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(o.Severity) { // not required
		return nil
	}

	// value enum
	if err := o.validateSeverityEnum("body"+"."+"severity", "body", *o.Severity); err != nil {
		return err
	}

	return nil
}

func (o *CreateAlertRuleBody) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(o.Filters) { // not required
		return nil
	}

	for i := 0; i < len(o.Filters); i++ {
		if swag.IsZero(o.Filters[i]) { // not required
			continue
		}

		if o.Filters[i] != nil {
			if err := o.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create alert rule body based on the context it is used
func (o *CreateAlertRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAlertRuleBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Params); i++ {
		if o.Params[i] != nil {
			if err := o.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *CreateAlertRuleBody) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Filters); i++ {
		if o.Filters[i] != nil {
			if err := o.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateAlertRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAlertRuleBody) UnmarshalBinary(b []byte) error {
	var res CreateAlertRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAlertRuleDefaultBody create alert rule default body
swagger:model CreateAlertRuleDefaultBody
*/
type CreateAlertRuleDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*CreateAlertRuleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this create alert rule default body
func (o *CreateAlertRuleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAlertRuleDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("CreateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CreateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create alert rule default body based on the context it is used
func (o *CreateAlertRuleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAlertRuleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CreateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CreateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateAlertRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAlertRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateAlertRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAlertRuleDefaultBodyDetailsItems0 create alert rule default body details items0
swagger:model CreateAlertRuleDefaultBodyDetailsItems0
*/
type CreateAlertRuleDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this create alert rule default body details items0
func (o *CreateAlertRuleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create alert rule default body details items0 based on context it is used
func (o *CreateAlertRuleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateAlertRuleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAlertRuleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CreateAlertRuleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAlertRuleOKBody create alert rule OK body
swagger:model CreateAlertRuleOKBody
*/
type CreateAlertRuleOKBody struct {
	// Rule ID.
	RuleID string `json:"rule_id,omitempty"`
}

// Validate validates this create alert rule OK body
func (o *CreateAlertRuleOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create alert rule OK body based on context it is used
func (o *CreateAlertRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateAlertRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAlertRuleOKBody) UnmarshalBinary(b []byte) error {
	var res CreateAlertRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAlertRuleParamsBodyFiltersItems0 Filter repsents a single filter condition.
swagger:model CreateAlertRuleParamsBodyFiltersItems0
*/
type CreateAlertRuleParamsBodyFiltersItems0 struct {
	// FilterType represents filter matching type.
	//
	//  - EQUAL: =
	//  - REGEX: =~
	// Enum: [FILTER_TYPE_INVALID EQUAL REGEX]
	Type *string `json:"type,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this create alert rule params body filters items0
func (o *CreateAlertRuleParamsBodyFiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createAlertRuleParamsBodyFiltersItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FILTER_TYPE_INVALID","EQUAL","REGEX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAlertRuleParamsBodyFiltersItems0TypeTypePropEnum = append(createAlertRuleParamsBodyFiltersItems0TypeTypePropEnum, v)
	}
}

const (

	// CreateAlertRuleParamsBodyFiltersItems0TypeFILTERTYPEINVALID captures enum value "FILTER_TYPE_INVALID"
	CreateAlertRuleParamsBodyFiltersItems0TypeFILTERTYPEINVALID string = "FILTER_TYPE_INVALID"

	// CreateAlertRuleParamsBodyFiltersItems0TypeEQUAL captures enum value "EQUAL"
	CreateAlertRuleParamsBodyFiltersItems0TypeEQUAL string = "EQUAL"

	// CreateAlertRuleParamsBodyFiltersItems0TypeREGEX captures enum value "REGEX"
	CreateAlertRuleParamsBodyFiltersItems0TypeREGEX string = "REGEX"
)

// prop value enum
func (o *CreateAlertRuleParamsBodyFiltersItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAlertRuleParamsBodyFiltersItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateAlertRuleParamsBodyFiltersItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create alert rule params body filters items0 based on context it is used
func (o *CreateAlertRuleParamsBodyFiltersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateAlertRuleParamsBodyFiltersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAlertRuleParamsBodyFiltersItems0) UnmarshalBinary(b []byte) error {
	var res CreateAlertRuleParamsBodyFiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAlertRuleParamsBodyParamsItems0 ParamValue represents a single rule parameter value for List, Change and Update APIs.
swagger:model CreateAlertRuleParamsBodyParamsItems0
*/
type CreateAlertRuleParamsBodyParamsItems0 struct {
	// Machine-readable name (ID) that is used in expression.
	Name string `json:"name,omitempty"`

	// ParamType represents template parameter type.
	// Enum: [PARAM_TYPE_INVALID BOOL FLOAT STRING]
	Type *string `json:"type,omitempty"`

	// Bool value.
	Bool bool `json:"bool,omitempty"`

	// Float value.
	Float float64 `json:"float,omitempty"`

	// String value.
	String string `json:"string,omitempty"`
}

// Validate validates this create alert rule params body params items0
func (o *CreateAlertRuleParamsBodyParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createAlertRuleParamsBodyParamsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARAM_TYPE_INVALID","BOOL","FLOAT","STRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAlertRuleParamsBodyParamsItems0TypeTypePropEnum = append(createAlertRuleParamsBodyParamsItems0TypeTypePropEnum, v)
	}
}

const (

	// CreateAlertRuleParamsBodyParamsItems0TypePARAMTYPEINVALID captures enum value "PARAM_TYPE_INVALID"
	CreateAlertRuleParamsBodyParamsItems0TypePARAMTYPEINVALID string = "PARAM_TYPE_INVALID"

	// CreateAlertRuleParamsBodyParamsItems0TypeBOOL captures enum value "BOOL"
	CreateAlertRuleParamsBodyParamsItems0TypeBOOL string = "BOOL"

	// CreateAlertRuleParamsBodyParamsItems0TypeFLOAT captures enum value "FLOAT"
	CreateAlertRuleParamsBodyParamsItems0TypeFLOAT string = "FLOAT"

	// CreateAlertRuleParamsBodyParamsItems0TypeSTRING captures enum value "STRING"
	CreateAlertRuleParamsBodyParamsItems0TypeSTRING string = "STRING"
)

// prop value enum
func (o *CreateAlertRuleParamsBodyParamsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAlertRuleParamsBodyParamsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateAlertRuleParamsBodyParamsItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create alert rule params body params items0 based on context it is used
func (o *CreateAlertRuleParamsBodyParamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateAlertRuleParamsBodyParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAlertRuleParamsBodyParamsItems0) UnmarshalBinary(b []byte) error {
	var res CreateAlertRuleParamsBodyParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
