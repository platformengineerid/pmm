// Code generated by go-swagger; DO NOT EDIT.

package proxy_sql

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

// AddProxySQLReader is a Reader for the AddProxySQL structure.
type AddProxySQLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProxySQLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProxySQLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddProxySQLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddProxySQLOK creates a AddProxySQLOK with default headers values
func NewAddProxySQLOK() *AddProxySQLOK {
	return &AddProxySQLOK{}
}

/*
AddProxySQLOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddProxySQLOK struct {
	Payload *AddProxySQLOKBody
}

func (o *AddProxySQLOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ProxySQL/Add][%d] addProxySqlOk  %+v", 200, o.Payload)
}

func (o *AddProxySQLOK) GetPayload() *AddProxySQLOKBody {
	return o.Payload
}

func (o *AddProxySQLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddProxySQLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProxySQLDefault creates a AddProxySQLDefault with default headers values
func NewAddProxySQLDefault(code int) *AddProxySQLDefault {
	return &AddProxySQLDefault{
		_statusCode: code,
	}
}

/*
AddProxySQLDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddProxySQLDefault struct {
	_statusCode int

	Payload *AddProxySQLDefaultBody
}

// Code gets the status code for the add proxy SQL default response
func (o *AddProxySQLDefault) Code() int {
	return o._statusCode
}

func (o *AddProxySQLDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ProxySQL/Add][%d] AddProxySQL default  %+v", o._statusCode, o.Payload)
}

func (o *AddProxySQLDefault) GetPayload() *AddProxySQLDefaultBody {
	return o.Payload
}

func (o *AddProxySQLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddProxySQLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddProxySQLBody add proxy SQL body
swagger:model AddProxySQLBody
*/
type AddProxySQLBody struct {
	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"node_id,omitempty"`

	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node and Service access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Service Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Service Access socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// ProxySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MetricsMode defines desired metrics mode for agent,
	// it can be pull, push or auto mode chosen by server.
	//
	//  - METRICS_MODE_UNSPECIFIED: Auto
	// Enum: [METRICS_MODE_UNSPECIFIED METRICS_MODE_PULL METRICS_MODE_PUSH]
	MetricsMode *string `json:"metrics_mode,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`

	// Log level for exporters
	//
	// - LOG_LEVEL_UNSPECIFIED: Auto
	// Enum: [LOG_LEVEL_UNSPECIFIED LOG_LEVEL_FATAL LOG_LEVEL_ERROR LOG_LEVEL_WARN LOG_LEVEL_INFO LOG_LEVEL_DEBUG]
	LogLevel *string `json:"log_level,omitempty"`

	// add node
	AddNode *AddProxySQLParamsBodyAddNode `json:"add_node,omitempty"`
}

// Validate validates this add proxy SQL body
func (o *AddProxySQLBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetricsMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlBodyTypeMetricsModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["METRICS_MODE_UNSPECIFIED","METRICS_MODE_PULL","METRICS_MODE_PUSH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlBodyTypeMetricsModePropEnum = append(addProxySqlBodyTypeMetricsModePropEnum, v)
	}
}

const (

	// AddProxySQLBodyMetricsModeMETRICSMODEUNSPECIFIED captures enum value "METRICS_MODE_UNSPECIFIED"
	AddProxySQLBodyMetricsModeMETRICSMODEUNSPECIFIED string = "METRICS_MODE_UNSPECIFIED"

	// AddProxySQLBodyMetricsModeMETRICSMODEPULL captures enum value "METRICS_MODE_PULL"
	AddProxySQLBodyMetricsModeMETRICSMODEPULL string = "METRICS_MODE_PULL"

	// AddProxySQLBodyMetricsModeMETRICSMODEPUSH captures enum value "METRICS_MODE_PUSH"
	AddProxySQLBodyMetricsModeMETRICSMODEPUSH string = "METRICS_MODE_PUSH"
)

// prop value enum
func (o *AddProxySQLBody) validateMetricsModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlBodyTypeMetricsModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLBody) validateMetricsMode(formats strfmt.Registry) error {
	if swag.IsZero(o.MetricsMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateMetricsModeEnum("body"+"."+"metrics_mode", "body", *o.MetricsMode); err != nil {
		return err
	}

	return nil
}

var addProxySqlBodyTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOG_LEVEL_UNSPECIFIED","LOG_LEVEL_FATAL","LOG_LEVEL_ERROR","LOG_LEVEL_WARN","LOG_LEVEL_INFO","LOG_LEVEL_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlBodyTypeLogLevelPropEnum = append(addProxySqlBodyTypeLogLevelPropEnum, v)
	}
}

const (

	// AddProxySQLBodyLogLevelLOGLEVELUNSPECIFIED captures enum value "LOG_LEVEL_UNSPECIFIED"
	AddProxySQLBodyLogLevelLOGLEVELUNSPECIFIED string = "LOG_LEVEL_UNSPECIFIED"

	// AddProxySQLBodyLogLevelLOGLEVELFATAL captures enum value "LOG_LEVEL_FATAL"
	AddProxySQLBodyLogLevelLOGLEVELFATAL string = "LOG_LEVEL_FATAL"

	// AddProxySQLBodyLogLevelLOGLEVELERROR captures enum value "LOG_LEVEL_ERROR"
	AddProxySQLBodyLogLevelLOGLEVELERROR string = "LOG_LEVEL_ERROR"

	// AddProxySQLBodyLogLevelLOGLEVELWARN captures enum value "LOG_LEVEL_WARN"
	AddProxySQLBodyLogLevelLOGLEVELWARN string = "LOG_LEVEL_WARN"

	// AddProxySQLBodyLogLevelLOGLEVELINFO captures enum value "LOG_LEVEL_INFO"
	AddProxySQLBodyLogLevelLOGLEVELINFO string = "LOG_LEVEL_INFO"

	// AddProxySQLBodyLogLevelLOGLEVELDEBUG captures enum value "LOG_LEVEL_DEBUG"
	AddProxySQLBodyLogLevelLOGLEVELDEBUG string = "LOG_LEVEL_DEBUG"
)

// prop value enum
func (o *AddProxySQLBody) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlBodyTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLBody) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("body"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

func (o *AddProxySQLBody) validateAddNode(formats strfmt.Registry) error {
	if swag.IsZero(o.AddNode) { // not required
		return nil
	}

	if o.AddNode != nil {
		if err := o.AddNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "add_node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "add_node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add proxy SQL body based on the context it is used
func (o *AddProxySQLBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAddNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLBody) contextValidateAddNode(ctx context.Context, formats strfmt.Registry) error {
	if o.AddNode != nil {
		if err := o.AddNode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "add_node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "add_node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLDefaultBody add proxy SQL default body
swagger:model AddProxySQLDefaultBody
*/
type AddProxySQLDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddProxySQLDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add proxy SQL default body
func (o *AddProxySQLDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddProxySQL default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddProxySQL default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add proxy SQL default body based on the context it is used
func (o *AddProxySQLDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddProxySQL default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddProxySQL default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLDefaultBodyDetailsItems0 add proxy SQL default body details items0
swagger:model AddProxySQLDefaultBodyDetailsItems0
*/
type AddProxySQLDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add proxy SQL default body details items0
func (o *AddProxySQLDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add proxy SQL default body details items0 based on context it is used
func (o *AddProxySQLDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddProxySQLDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLOKBody add proxy SQL OK body
swagger:model AddProxySQLOKBody
*/
type AddProxySQLOKBody struct {
	// proxysql exporter
	ProxysqlExporter *AddProxySQLOKBodyProxysqlExporter `json:"proxysql_exporter,omitempty"`

	// service
	Service *AddProxySQLOKBodyService `json:"service,omitempty"`
}

// Validate validates this add proxy SQL OK body
func (o *AddProxySQLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysqlExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLOKBody) validateProxysqlExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.ProxysqlExporter) { // not required
		return nil
	}

	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlOk" + "." + "proxysql_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addProxySqlOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddProxySQLOKBody) validateService(formats strfmt.Registry) error {
	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlOk" + "." + "service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addProxySqlOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add proxy SQL OK body based on the context it is used
func (o *AddProxySQLOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProxysqlExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLOKBody) contextValidateProxysqlExporter(ctx context.Context, formats strfmt.Registry) error {
	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlOk" + "." + "proxysql_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addProxySqlOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddProxySQLOKBody) contextValidateService(ctx context.Context, formats strfmt.Registry) error {
	if o.Service != nil {
		if err := o.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlOk" + "." + "service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addProxySqlOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLOKBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLOKBodyProxysqlExporter ProxySQLExporter runs on Generic or Container Node and exposes ProxySQL Service metrics.
swagger:model AddProxySQLOKBodyProxysqlExporter
*/
type AddProxySQLOKBodyProxysqlExporter struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	//
	// Status fields below.
	DisabledCollectors []string `json:"disabled_collectors"`

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

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// Log level for exporters
	//
	// - LOG_LEVEL_UNSPECIFIED: Auto
	// Enum: [LOG_LEVEL_UNSPECIFIED LOG_LEVEL_FATAL LOG_LEVEL_ERROR LOG_LEVEL_WARN LOG_LEVEL_INFO LOG_LEVEL_DEBUG]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this add proxy SQL OK body proxysql exporter
func (o *AddProxySQLOKBodyProxysqlExporter) Validate(formats strfmt.Registry) error {
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

var addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_UNSPECIFIED","AGENT_STATUS_STARTING","AGENT_STATUS_RUNNING","AGENT_STATUS_WAITING","AGENT_STATUS_STOPPING","AGENT_STATUS_DONE","AGENT_STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum = append(addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSUNSPECIFIED captures enum value "AGENT_STATUS_UNSPECIFIED"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSUNSPECIFIED string = "AGENT_STATUS_UNSPECIFIED"

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSSTARTING captures enum value "AGENT_STATUS_STARTING"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSSTARTING string = "AGENT_STATUS_STARTING"

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSRUNNING captures enum value "AGENT_STATUS_RUNNING"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSRUNNING string = "AGENT_STATUS_RUNNING"

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSWAITING captures enum value "AGENT_STATUS_WAITING"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSWAITING string = "AGENT_STATUS_WAITING"

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSSTOPPING captures enum value "AGENT_STATUS_STOPPING"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSSTOPPING string = "AGENT_STATUS_STOPPING"

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSDONE captures enum value "AGENT_STATUS_DONE"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSDONE string = "AGENT_STATUS_DONE"

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSUNKNOWN captures enum value "AGENT_STATUS_UNKNOWN"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSUNKNOWN string = "AGENT_STATUS_UNKNOWN"
)

// prop value enum
func (o *AddProxySQLOKBodyProxysqlExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLOKBodyProxysqlExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addProxySqlOk"+"."+"proxysql_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var addProxySqlOkBodyProxysqlExporterTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOG_LEVEL_UNSPECIFIED","LOG_LEVEL_FATAL","LOG_LEVEL_ERROR","LOG_LEVEL_WARN","LOG_LEVEL_INFO","LOG_LEVEL_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlOkBodyProxysqlExporterTypeLogLevelPropEnum = append(addProxySqlOkBodyProxysqlExporterTypeLogLevelPropEnum, v)
	}
}

const (

	// AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELUNSPECIFIED captures enum value "LOG_LEVEL_UNSPECIFIED"
	AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELUNSPECIFIED string = "LOG_LEVEL_UNSPECIFIED"

	// AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELFATAL captures enum value "LOG_LEVEL_FATAL"
	AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELFATAL string = "LOG_LEVEL_FATAL"

	// AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELERROR captures enum value "LOG_LEVEL_ERROR"
	AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELERROR string = "LOG_LEVEL_ERROR"

	// AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELWARN captures enum value "LOG_LEVEL_WARN"
	AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELWARN string = "LOG_LEVEL_WARN"

	// AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELINFO captures enum value "LOG_LEVEL_INFO"
	AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELINFO string = "LOG_LEVEL_INFO"

	// AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELDEBUG captures enum value "LOG_LEVEL_DEBUG"
	AddProxySQLOKBodyProxysqlExporterLogLevelLOGLEVELDEBUG string = "LOG_LEVEL_DEBUG"
)

// prop value enum
func (o *AddProxySQLOKBodyProxysqlExporter) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlOkBodyProxysqlExporterTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLOKBodyProxysqlExporter) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("addProxySqlOk"+"."+"proxysql_exporter"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add proxy SQL OK body proxysql exporter based on context it is used
func (o *AddProxySQLOKBodyProxysqlExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLOKBodyProxysqlExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLOKBodyProxysqlExporter) UnmarshalBinary(b []byte) error {
	var res AddProxySQLOKBodyProxysqlExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLOKBodyService ProxySQLService represents a generic ProxySQL instance.
swagger:model AddProxySQLOKBodyService
*/
type AddProxySQLOKBodyService struct {
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

	// ProxySQL version.
	Version string `json:"version,omitempty"`
}

// Validate validates this add proxy SQL OK body service
func (o *AddProxySQLOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add proxy SQL OK body service based on context it is used
func (o *AddProxySQLOKBodyService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddProxySQLOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLParamsBodyAddNode AddNodeParams holds node params and is used to add new node to inventory while adding new service.
swagger:model AddProxySQLParamsBodyAddNode
*/
type AddProxySQLParamsBodyAddNode struct {
	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_UNSPECIFIED NODE_TYPE_GENERIC_NODE NODE_TYPE_CONTAINER_NODE NODE_TYPE_REMOTE_NODE NODE_TYPE_REMOTE_RDS_NODE NODE_TYPE_REMOTE_AZURE_DATABASE_NODE]
	NodeType *string `json:"node_type,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels for Node.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add proxy SQL params body add node
func (o *AddProxySQLParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_UNSPECIFIED","NODE_TYPE_GENERIC_NODE","NODE_TYPE_CONTAINER_NODE","NODE_TYPE_REMOTE_NODE","NODE_TYPE_REMOTE_RDS_NODE","NODE_TYPE_REMOTE_AZURE_DATABASE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum = append(addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEUNSPECIFIED captures enum value "NODE_TYPE_UNSPECIFIED"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEUNSPECIFIED string = "NODE_TYPE_UNSPECIFIED"

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEGENERICNODE captures enum value "NODE_TYPE_GENERIC_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEGENERICNODE string = "NODE_TYPE_GENERIC_NODE"

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPECONTAINERNODE captures enum value "NODE_TYPE_CONTAINER_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPECONTAINERNODE string = "NODE_TYPE_CONTAINER_NODE"

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEREMOTENODE captures enum value "NODE_TYPE_REMOTE_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEREMOTENODE string = "NODE_TYPE_REMOTE_NODE"

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEREMOTERDSNODE captures enum value "NODE_TYPE_REMOTE_RDS_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEREMOTERDSNODE string = "NODE_TYPE_REMOTE_RDS_NODE"

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEREMOTEAZUREDATABASENODE captures enum value "NODE_TYPE_REMOTE_AZURE_DATABASE_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEREMOTEAZUREDATABASENODE string = "NODE_TYPE_REMOTE_AZURE_DATABASE_NODE"
)

// prop value enum
func (o *AddProxySQLParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"add_node"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add proxy SQL params body add node based on context it is used
func (o *AddProxySQLParamsBodyAddNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddProxySQLParamsBodyAddNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
