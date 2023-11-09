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

// ChangeQANMySQLSlowlogAgentReader is a Reader for the ChangeQANMySQLSlowlogAgent structure.
type ChangeQANMySQLSlowlogAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeQANMySQLSlowlogAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeQANMySQLSlowlogAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeQANMySQLSlowlogAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeQANMySQLSlowlogAgentOK creates a ChangeQANMySQLSlowlogAgentOK with default headers values
func NewChangeQANMySQLSlowlogAgentOK() *ChangeQANMySQLSlowlogAgentOK {
	return &ChangeQANMySQLSlowlogAgentOK{}
}

/*
ChangeQANMySQLSlowlogAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeQANMySQLSlowlogAgentOK struct {
	Payload *ChangeQANMySQLSlowlogAgentOKBody
}

func (o *ChangeQANMySQLSlowlogAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANMySQLSlowlogAgent][%d] changeQanMySqlSlowlogAgentOk  %+v", 200, o.Payload)
}

func (o *ChangeQANMySQLSlowlogAgentOK) GetPayload() *ChangeQANMySQLSlowlogAgentOKBody {
	return o.Payload
}

func (o *ChangeQANMySQLSlowlogAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeQANMySQLSlowlogAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeQANMySQLSlowlogAgentDefault creates a ChangeQANMySQLSlowlogAgentDefault with default headers values
func NewChangeQANMySQLSlowlogAgentDefault(code int) *ChangeQANMySQLSlowlogAgentDefault {
	return &ChangeQANMySQLSlowlogAgentDefault{
		_statusCode: code,
	}
}

/*
ChangeQANMySQLSlowlogAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeQANMySQLSlowlogAgentDefault struct {
	_statusCode int

	Payload *ChangeQANMySQLSlowlogAgentDefaultBody
}

// Code gets the status code for the change QAN my SQL slowlog agent default response
func (o *ChangeQANMySQLSlowlogAgentDefault) Code() int {
	return o._statusCode
}

func (o *ChangeQANMySQLSlowlogAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANMySQLSlowlogAgent][%d] ChangeQANMySQLSlowlogAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeQANMySQLSlowlogAgentDefault) GetPayload() *ChangeQANMySQLSlowlogAgentDefaultBody {
	return o.Payload
}

func (o *ChangeQANMySQLSlowlogAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeQANMySQLSlowlogAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ChangeQANMySQLSlowlogAgentBody change QAN my SQL slowlog agent body
swagger:model ChangeQANMySQLSlowlogAgentBody
*/
type ChangeQANMySQLSlowlogAgentBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeQANMySQLSlowlogAgentParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change QAN my SQL slowlog agent body
func (o *ChangeQANMySQLSlowlogAgentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentBody) validateCommon(formats strfmt.Registry) error {
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

// ContextValidate validate this change QAN my SQL slowlog agent body based on the context it is used
func (o *ChangeQANMySQLSlowlogAgentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentBody) contextValidateCommon(ctx context.Context, formats strfmt.Registry) error {
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
func (o *ChangeQANMySQLSlowlogAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLSlowlogAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANMySQLSlowlogAgentDefaultBody change QAN my SQL slowlog agent default body
swagger:model ChangeQANMySQLSlowlogAgentDefaultBody
*/
type ChangeQANMySQLSlowlogAgentDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change QAN my SQL slowlog agent default body
func (o *ChangeQANMySQLSlowlogAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeQANMySQLSlowlogAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeQANMySQLSlowlogAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change QAN my SQL slowlog agent default body based on the context it is used
func (o *ChangeQANMySQLSlowlogAgentDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeQANMySQLSlowlogAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeQANMySQLSlowlogAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLSlowlogAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0 change QAN my SQL slowlog agent default body details items0
swagger:model ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0
*/
type ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this change QAN my SQL slowlog agent default body details items0
func (o *ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change QAN my SQL slowlog agent default body details items0 based on context it is used
func (o *ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLSlowlogAgentDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANMySQLSlowlogAgentOKBody change QAN my SQL slowlog agent OK body
swagger:model ChangeQANMySQLSlowlogAgentOKBody
*/
type ChangeQANMySQLSlowlogAgentOKBody struct {
	// qan mysql slowlog agent
	QANMysqlSlowlogAgent *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent `json:"qan_mysql_slowlog_agent,omitempty"`
}

// Validate validates this change QAN my SQL slowlog agent OK body
func (o *ChangeQANMySQLSlowlogAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANMysqlSlowlogAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentOKBody) validateQANMysqlSlowlogAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.QANMysqlSlowlogAgent) { // not required
		return nil
	}

	if o.QANMysqlSlowlogAgent != nil {
		if err := o.QANMysqlSlowlogAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanMySqlSlowlogAgentOk" + "." + "qan_mysql_slowlog_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeQanMySqlSlowlogAgentOk" + "." + "qan_mysql_slowlog_agent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change QAN my SQL slowlog agent OK body based on the context it is used
func (o *ChangeQANMySQLSlowlogAgentOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateQANMysqlSlowlogAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentOKBody) contextValidateQANMysqlSlowlogAgent(ctx context.Context, formats strfmt.Registry) error {
	if o.QANMysqlSlowlogAgent != nil {
		if err := o.QANMysqlSlowlogAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanMySqlSlowlogAgentOk" + "." + "qan_mysql_slowlog_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeQanMySqlSlowlogAgentOk" + "." + "qan_mysql_slowlog_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLSlowlogAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent QANMySQLSlowlogAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent
*/
type ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Client certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// Password for decrypting tls_cert.
	TLSKey string `json:"tls_key,omitempty"`

	// Disable parsing comments from queries and showing them in QAN.
	DisableCommentsParsing bool `json:"disable_comments_parsing,omitempty"`

	// Limit query length in QAN (default: server-defined; -1: no limit)
	MaxQueryLength int32 `json:"max_query_length,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// Slowlog file is rotated at this size if > 0.
	MaxSlowlogFileSize string `json:"max_slowlog_file_size,omitempty"`

	// Custom user-assigned labels.
	//
	// Status fields below.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_UNSPECIFIED STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// mod tidy
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// Log level for exporters
	// Enum: [auto fatal error warn info debug]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this change QAN my SQL slowlog agent OK body QAN mysql slowlog agent
func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) Validate(formats strfmt.Registry) error {
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

var changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_UNSPECIFIED","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum = append(changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum, v)
	}
}

const (

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusAGENTSTATUSUNSPECIFIED captures enum value "AGENT_STATUS_UNSPECIFIED"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusAGENTSTATUSUNSPECIFIED string = "AGENT_STATUS_UNSPECIFIED"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTARTING captures enum value "STARTING"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTARTING string = "STARTING"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusRUNNING captures enum value "RUNNING"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusRUNNING string = "RUNNING"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusWAITING captures enum value "WAITING"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusWAITING string = "WAITING"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTOPPING captures enum value "STOPPING"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTOPPING string = "STOPPING"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusDONE captures enum value "DONE"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusDONE string = "DONE"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusUNKNOWN captures enum value "UNKNOWN"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeQanMySqlSlowlogAgentOk"+"."+"qan_mysql_slowlog_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeLogLevelPropEnum = append(changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeLogLevelPropEnum, v)
	}
}

const (

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelAuto captures enum value "auto"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelAuto string = "auto"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelFatal captures enum value "fatal"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelFatal string = "fatal"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelError captures enum value "error"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelError string = "error"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelWarn captures enum value "warn"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelWarn string = "warn"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelInfo captures enum value "info"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelInfo string = "info"

	// ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelDebug captures enum value "debug"
	ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentLogLevelDebug string = "debug"
)

// prop value enum
func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("changeQanMySqlSlowlogAgentOk"+"."+"qan_mysql_slowlog_agent"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change QAN my SQL slowlog agent OK body QAN mysql slowlog agent based on context it is used
func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeQANMySQLSlowlogAgentParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeQANMySQLSlowlogAgentParamsBodyCommon
*/
type ChangeQANMySQLSlowlogAgentParamsBodyCommon struct {
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

// Validate validates this change QAN my SQL slowlog agent params body common
func (o *ChangeQANMySQLSlowlogAgentParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change QAN my SQL slowlog agent params body common based on context it is used
func (o *ChangeQANMySQLSlowlogAgentParamsBodyCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLSlowlogAgentParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLSlowlogAgentParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
