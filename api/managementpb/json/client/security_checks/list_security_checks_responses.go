// Code generated by go-swagger; DO NOT EDIT.

package security_checks

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

// ListSecurityChecksReader is a Reader for the ListSecurityChecks structure.
type ListSecurityChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSecurityChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSecurityChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSecurityChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSecurityChecksOK creates a ListSecurityChecksOK with default headers values
func NewListSecurityChecksOK() *ListSecurityChecksOK {
	return &ListSecurityChecksOK{}
}

/*
ListSecurityChecksOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListSecurityChecksOK struct {
	Payload *ListSecurityChecksOKBody
}

func (o *ListSecurityChecksOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/List][%d] listSecurityChecksOk  %+v", 200, o.Payload)
}

func (o *ListSecurityChecksOK) GetPayload() *ListSecurityChecksOKBody {
	return o.Payload
}

func (o *ListSecurityChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListSecurityChecksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSecurityChecksDefault creates a ListSecurityChecksDefault with default headers values
func NewListSecurityChecksDefault(code int) *ListSecurityChecksDefault {
	return &ListSecurityChecksDefault{
		_statusCode: code,
	}
}

/*
ListSecurityChecksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListSecurityChecksDefault struct {
	_statusCode int

	Payload *ListSecurityChecksDefaultBody
}

// Code gets the status code for the list security checks default response
func (o *ListSecurityChecksDefault) Code() int {
	return o._statusCode
}

func (o *ListSecurityChecksDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/List][%d] ListSecurityChecks default  %+v", o._statusCode, o.Payload)
}

func (o *ListSecurityChecksDefault) GetPayload() *ListSecurityChecksDefaultBody {
	return o.Payload
}

func (o *ListSecurityChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListSecurityChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListSecurityChecksDefaultBody list security checks default body
swagger:model ListSecurityChecksDefaultBody
*/
type ListSecurityChecksDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListSecurityChecksDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list security checks default body
func (o *ListSecurityChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list security checks default body based on the context it is used
func (o *ListSecurityChecksDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListSecurityChecksDefaultBodyDetailsItems0 list security checks default body details items0
swagger:model ListSecurityChecksDefaultBodyDetailsItems0
*/
type ListSecurityChecksDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list security checks default body details items0
func (o *ListSecurityChecksDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list security checks default body details items0 based on context it is used
func (o *ListSecurityChecksDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListSecurityChecksOKBody list security checks OK body
swagger:model ListSecurityChecksOKBody
*/
type ListSecurityChecksOKBody struct {
	// checks
	Checks []*ListSecurityChecksOKBodyChecksItems0 `json:"checks"`
}

// Validate validates this list security checks OK body
func (o *ListSecurityChecksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksOKBody) validateChecks(formats strfmt.Registry) error {
	if swag.IsZero(o.Checks) { // not required
		return nil
	}

	for i := 0; i < len(o.Checks); i++ {
		if swag.IsZero(o.Checks[i]) { // not required
			continue
		}

		if o.Checks[i] != nil {
			if err := o.Checks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listSecurityChecksOk" + "." + "checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listSecurityChecksOk" + "." + "checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list security checks OK body based on the context it is used
func (o *ListSecurityChecksOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksOKBody) contextValidateChecks(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Checks); i++ {
		if o.Checks[i] != nil {
			if err := o.Checks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listSecurityChecksOk" + "." + "checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listSecurityChecksOk" + "." + "checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBody) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListSecurityChecksOKBodyChecksItems0 SecurityCheck contains check name and status.
swagger:model ListSecurityChecksOKBodyChecksItems0
*/
type ListSecurityChecksOKBodyChecksItems0 struct {
	// Machine-readable name (ID) that is used in expression.
	Name string `json:"name,omitempty"`

	// True if that check is disabled.
	Disabled bool `json:"disabled,omitempty"`

	// Long human-readable description.
	Description string `json:"description,omitempty"`

	// Short human-readable summary.
	Summary string `json:"summary,omitempty"`

	// SecurityCheckInterval represents possible execution interval values for checks.
	// Enum: [SECURITY_CHECK_INTERVAL_UNSPECIFIED SECURITY_CHECK_INTERVAL_STANDARD SECURITY_CHECK_INTERVAL_FREQUENT SECURITY_CHECK_INTERVAL_RARE]
	Interval *string `json:"interval,omitempty"`

	// family
	// Enum: [ADVISOR_CHECK_FAMILY_UNSPECIFIED ADVISOR_CHECK_FAMILY_MYSQL ADVISOR_CHECK_FAMILY_POSTGRESQL ADVISOR_CHECK_FAMILY_MONGODB]
	Family *string `json:"family,omitempty"`
}

// Validate validates this list security checks OK body checks items0
func (o *ListSecurityChecksOKBodyChecksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listSecurityChecksOkBodyChecksItems0TypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SECURITY_CHECK_INTERVAL_UNSPECIFIED","SECURITY_CHECK_INTERVAL_STANDARD","SECURITY_CHECK_INTERVAL_FREQUENT","SECURITY_CHECK_INTERVAL_RARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listSecurityChecksOkBodyChecksItems0TypeIntervalPropEnum = append(listSecurityChecksOkBodyChecksItems0TypeIntervalPropEnum, v)
	}
}

const (

	// ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALUNSPECIFIED captures enum value "SECURITY_CHECK_INTERVAL_UNSPECIFIED"
	ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALUNSPECIFIED string = "SECURITY_CHECK_INTERVAL_UNSPECIFIED"

	// ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALSTANDARD captures enum value "SECURITY_CHECK_INTERVAL_STANDARD"
	ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALSTANDARD string = "SECURITY_CHECK_INTERVAL_STANDARD"

	// ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALFREQUENT captures enum value "SECURITY_CHECK_INTERVAL_FREQUENT"
	ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALFREQUENT string = "SECURITY_CHECK_INTERVAL_FREQUENT"

	// ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALRARE captures enum value "SECURITY_CHECK_INTERVAL_RARE"
	ListSecurityChecksOKBodyChecksItems0IntervalSECURITYCHECKINTERVALRARE string = "SECURITY_CHECK_INTERVAL_RARE"
)

// prop value enum
func (o *ListSecurityChecksOKBodyChecksItems0) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listSecurityChecksOkBodyChecksItems0TypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListSecurityChecksOKBodyChecksItems0) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(o.Interval) { // not required
		return nil
	}

	// value enum
	if err := o.validateIntervalEnum("interval", "body", *o.Interval); err != nil {
		return err
	}

	return nil
}

var listSecurityChecksOkBodyChecksItems0TypeFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADVISOR_CHECK_FAMILY_UNSPECIFIED","ADVISOR_CHECK_FAMILY_MYSQL","ADVISOR_CHECK_FAMILY_POSTGRESQL","ADVISOR_CHECK_FAMILY_MONGODB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listSecurityChecksOkBodyChecksItems0TypeFamilyPropEnum = append(listSecurityChecksOkBodyChecksItems0TypeFamilyPropEnum, v)
	}
}

const (

	// ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYUNSPECIFIED captures enum value "ADVISOR_CHECK_FAMILY_UNSPECIFIED"
	ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYUNSPECIFIED string = "ADVISOR_CHECK_FAMILY_UNSPECIFIED"

	// ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYMYSQL captures enum value "ADVISOR_CHECK_FAMILY_MYSQL"
	ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYMYSQL string = "ADVISOR_CHECK_FAMILY_MYSQL"

	// ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYPOSTGRESQL captures enum value "ADVISOR_CHECK_FAMILY_POSTGRESQL"
	ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYPOSTGRESQL string = "ADVISOR_CHECK_FAMILY_POSTGRESQL"

	// ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYMONGODB captures enum value "ADVISOR_CHECK_FAMILY_MONGODB"
	ListSecurityChecksOKBodyChecksItems0FamilyADVISORCHECKFAMILYMONGODB string = "ADVISOR_CHECK_FAMILY_MONGODB"
)

// prop value enum
func (o *ListSecurityChecksOKBodyChecksItems0) validateFamilyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listSecurityChecksOkBodyChecksItems0TypeFamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListSecurityChecksOKBodyChecksItems0) validateFamily(formats strfmt.Registry) error {
	if swag.IsZero(o.Family) { // not required
		return nil
	}

	// value enum
	if err := o.validateFamilyEnum("family", "body", *o.Family); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list security checks OK body checks items0 based on context it is used
func (o *ListSecurityChecksOKBodyChecksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyChecksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyChecksItems0) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyChecksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
