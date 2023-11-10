// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// ChangeQANPostgreSQLPgStatMonitorAgentReader is a Reader for the ChangeQANPostgreSQLPgStatMonitorAgent structure.
type ChangeQANPostgreSQLPgStatMonitorAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeQANPostgreSQLPgStatMonitorAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeQANPostgreSQLPgStatMonitorAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeQANPostgreSQLPgStatMonitorAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeQANPostgreSQLPgStatMonitorAgentOK creates a ChangeQANPostgreSQLPgStatMonitorAgentOK with default headers values
func NewChangeQANPostgreSQLPgStatMonitorAgentOK() *ChangeQANPostgreSQLPgStatMonitorAgentOK {
	return &ChangeQANPostgreSQLPgStatMonitorAgentOK{}
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeQANPostgreSQLPgStatMonitorAgentOK struct {
	Payload *ChangeQANPostgreSQLPgStatMonitorAgentOKBody
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANPostgreSQLPgStatMonitorAgent][%d] changeQanPostgreSqlPgStatMonitorAgentOk  %+v", 200, o.Payload)
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOK) GetPayload() *ChangeQANPostgreSQLPgStatMonitorAgentOKBody {
	return o.Payload
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeQANPostgreSQLPgStatMonitorAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeQANPostgreSQLPgStatMonitorAgentDefault creates a ChangeQANPostgreSQLPgStatMonitorAgentDefault with default headers values
func NewChangeQANPostgreSQLPgStatMonitorAgentDefault(code int) *ChangeQANPostgreSQLPgStatMonitorAgentDefault {
	return &ChangeQANPostgreSQLPgStatMonitorAgentDefault{
		_statusCode: code,
	}
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeQANPostgreSQLPgStatMonitorAgentDefault struct {
	_statusCode int

	Payload *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody
}

// Code gets the status code for the change QAN postgre SQL pg stat monitor agent default response
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefault) Code() int {
	return o._statusCode
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANPostgreSQLPgStatMonitorAgent][%d] ChangeQANPostgreSQLPgStatMonitorAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefault) GetPayload() *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody {
	return o.Payload
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentBody change QAN postgre SQL pg stat monitor agent body
swagger:model ChangeQANPostgreSQLPgStatMonitorAgentBody
*/
type ChangeQANPostgreSQLPgStatMonitorAgentBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change QAN postgre SQL pg stat monitor agent body
func (o *ChangeQANPostgreSQLPgStatMonitorAgentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentBody) validateCommon(formats strfmt.Registry) error {
	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change QAN postgre SQL pg stat monitor agent body based on the context it is used
func (o *ChangeQANPostgreSQLPgStatMonitorAgentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentBody) contextValidateCommon(ctx context.Context, formats strfmt.Registry) error {
	if o.Common != nil {
		if err := o.Common.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatMonitorAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody change QAN postgre SQL pg stat monitor agent default body
swagger:model ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody
*/
type ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change QAN postgre SQL pg stat monitor agent default body
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeQANPostgreSQLPgStatMonitorAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeQANPostgreSQLPgStatMonitorAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change QAN postgre SQL pg stat monitor agent default body based on the context it is used
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeQANPostgreSQLPgStatMonitorAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeQANPostgreSQLPgStatMonitorAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatMonitorAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0 change QAN postgre SQL pg stat monitor agent default body details items0
swagger:model ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0
*/
type ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this change QAN postgre SQL pg stat monitor agent default body details items0
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change QAN postgre SQL pg stat monitor agent default body details items0 based on context it is used
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatMonitorAgentDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentOKBody change QAN postgre SQL pg stat monitor agent OK body
swagger:model ChangeQANPostgreSQLPgStatMonitorAgentOKBody
*/
type ChangeQANPostgreSQLPgStatMonitorAgentOKBody struct {
	// qan postgresql pgstatmonitor agent
	QANPostgresqlPgstatmonitorAgent *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent `json:"qan_postgresql_pgstatmonitor_agent,omitempty"`
}

// Validate validates this change QAN postgre SQL pg stat monitor agent OK body
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANPostgresqlPgstatmonitorAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBody) validateQANPostgresqlPgstatmonitorAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.QANPostgresqlPgstatmonitorAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatmonitorAgent != nil {
		if err := o.QANPostgresqlPgstatmonitorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanPostgreSqlPgStatMonitorAgentOk" + "." + "qan_postgresql_pgstatmonitor_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeQanPostgreSqlPgStatMonitorAgentOk" + "." + "qan_postgresql_pgstatmonitor_agent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change QAN postgre SQL pg stat monitor agent OK body based on the context it is used
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateQANPostgresqlPgstatmonitorAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBody) contextValidateQANPostgresqlPgstatmonitorAgent(ctx context.Context, formats strfmt.Registry) error {
	if o.QANPostgresqlPgstatmonitorAgent != nil {
		if err := o.QANPostgresqlPgstatmonitorAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanPostgreSqlPgStatMonitorAgentOk" + "." + "qan_postgresql_pgstatmonitor_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeQanPostgreSqlPgStatMonitorAgentOk" + "." + "qan_postgresql_pgstatmonitor_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatMonitorAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent QANPostgreSQLPgStatMonitorAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent
*/
type ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for getting pg stat monitor data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Disable parsing comments from queries and showing them in QAN.
	DisableCommentsParsing bool `json:"disable_comments_parsing,omitempty"`

	// Limit query length in QAN (default: server-defined; -1: no limit).
	MaxQueryLength int32 `json:"max_query_length,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// Custom user-assigned labels.
	//
	// Status fields below.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - AGENT_STATUS_STARTING: Agent is starting.
	//  - AGENT_STATUS_RUNNING: Agent is running.
	//  - AGENT_STATUS_WAITING: Agent encountered error and will be restarted automatically soon.
	//  - AGENT_STATUS_STOPPING: Agent is stopping.
	//  - AGENT_STATUS_DONE: Agent finished.
	//  - AGENT_STATUS_UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_UNSPECIFIED AGENT_STATUS_STARTING AGENT_STATUS_RUNNING AGENT_STATUS_WAITING AGENT_STATUS_STOPPING AGENT_STATUS_DONE AGENT_STATUS_UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// Log level for exporters
	//
	// - LOG_LEVEL_UNSPECIFIED: Auto
	// Enum: [LOG_LEVEL_UNSPECIFIED LOG_LEVEL_FATAL LOG_LEVEL_ERROR LOG_LEVEL_WARN LOG_LEVEL_INFO LOG_LEVEL_DEBUG]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this change QAN postgre SQL pg stat monitor agent OK body QAN postgresql pgstatmonitor agent
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_UNSPECIFIED","AGENT_STATUS_STARTING","AGENT_STATUS_RUNNING","AGENT_STATUS_WAITING","AGENT_STATUS_STOPPING","AGENT_STATUS_DONE","AGENT_STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum = append(changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum, v)
	}
}

const (

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSUNSPECIFIED captures enum value "AGENT_STATUS_UNSPECIFIED"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSUNSPECIFIED string = "AGENT_STATUS_UNSPECIFIED"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSSTARTING captures enum value "AGENT_STATUS_STARTING"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSSTARTING string = "AGENT_STATUS_STARTING"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSRUNNING captures enum value "AGENT_STATUS_RUNNING"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSRUNNING string = "AGENT_STATUS_RUNNING"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSWAITING captures enum value "AGENT_STATUS_WAITING"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSWAITING string = "AGENT_STATUS_WAITING"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSSTOPPING captures enum value "AGENT_STATUS_STOPPING"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSSTOPPING string = "AGENT_STATUS_STOPPING"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSDONE captures enum value "AGENT_STATUS_DONE"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSDONE string = "AGENT_STATUS_DONE"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSUNKNOWN captures enum value "AGENT_STATUS_UNKNOWN"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSUNKNOWN string = "AGENT_STATUS_UNKNOWN"
)

// prop value enum
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeQanPostgreSqlPgStatMonitorAgentOk"+"."+"qan_postgresql_pgstatmonitor_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOG_LEVEL_UNSPECIFIED","LOG_LEVEL_FATAL","LOG_LEVEL_ERROR","LOG_LEVEL_WARN","LOG_LEVEL_INFO","LOG_LEVEL_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeLogLevelPropEnum = append(changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeLogLevelPropEnum, v)
	}
}

const (

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELUNSPECIFIED captures enum value "LOG_LEVEL_UNSPECIFIED"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELUNSPECIFIED string = "LOG_LEVEL_UNSPECIFIED"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELFATAL captures enum value "LOG_LEVEL_FATAL"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELFATAL string = "LOG_LEVEL_FATAL"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELERROR captures enum value "LOG_LEVEL_ERROR"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELERROR string = "LOG_LEVEL_ERROR"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELWARN captures enum value "LOG_LEVEL_WARN"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELWARN string = "LOG_LEVEL_WARN"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELINFO captures enum value "LOG_LEVEL_INFO"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELINFO string = "LOG_LEVEL_INFO"

	// ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELDEBUG captures enum value "LOG_LEVEL_DEBUG"
	ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentLogLevelLOGLEVELDEBUG string = "LOG_LEVEL_DEBUG"
)

// prop value enum
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("changeQanPostgreSqlPgStatMonitorAgentOk"+"."+"qan_postgresql_pgstatmonitor_agent"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change QAN postgre SQL pg stat monitor agent OK body QAN postgresql pgstatmonitor agent based on context it is used
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon
*/
type ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon struct {
	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change QAN postgre SQL pg stat monitor agent params body common
func (o *ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change QAN postgre SQL pg stat monitor agent params body common based on context it is used
func (o *ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatMonitorAgentParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
