// Code generated by go-swagger; DO NOT EDIT.

package agent

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
	"github.com/go-openapi/validate"
)

// ListAgentsReader is a Reader for the ListAgents structure.
type ListAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAgentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAgentsOK creates a ListAgentsOK with default headers values
func NewListAgentsOK() *ListAgentsOK {
	return &ListAgentsOK{}
}

/*
ListAgentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListAgentsOK struct {
	Payload *ListAgentsOKBody
}

func (o *ListAgentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Agent/List][%d] listAgentsOk  %+v", 200, o.Payload)
}

func (o *ListAgentsOK) GetPayload() *ListAgentsOKBody {
	return o.Payload
}

func (o *ListAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListAgentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAgentsDefault creates a ListAgentsDefault with default headers values
func NewListAgentsDefault(code int) *ListAgentsDefault {
	return &ListAgentsDefault{
		_statusCode: code,
	}
}

/*
ListAgentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListAgentsDefault struct {
	_statusCode int

	Payload *ListAgentsDefaultBody
}

// Code gets the status code for the list agents default response
func (o *ListAgentsDefault) Code() int {
	return o._statusCode
}

func (o *ListAgentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Agent/List][%d] ListAgents default  %+v", o._statusCode, o.Payload)
}

func (o *ListAgentsDefault) GetPayload() *ListAgentsDefaultBody {
	return o.Payload
}

func (o *ListAgentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListAgentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListAgentsBody Only one of the parameters below must be set.
swagger:model ListAgentsBody
*/
type ListAgentsBody struct {
	// Return only Agents that relate to a specific ServiceID.
	ServiceID string `json:"service_id,omitempty"`

	// Return only Agents that relate to a specific NodeID.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this list agents body
func (o *ListAgentsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list agents body based on context it is used
func (o *ListAgentsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsBody) UnmarshalBinary(b []byte) error {
	var res ListAgentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsDefaultBody list agents default body
swagger:model ListAgentsDefaultBody
*/
type ListAgentsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListAgentsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list agents default body
func (o *ListAgentsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListAgents default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListAgents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list agents default body based on the context it is used
func (o *ListAgentsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListAgents default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListAgents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListAgentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsDefaultBodyDetailsItems0 list agents default body details items0
swagger:model ListAgentsDefaultBodyDetailsItems0
*/
type ListAgentsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list agents default body details items0
func (o *ListAgentsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list agents default body details items0 based on context it is used
func (o *ListAgentsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListAgentsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsOKBody list agents OK body
swagger:model ListAgentsOKBody
*/
type ListAgentsOKBody struct {
	// List of Agents.
	Agents []*ListAgentsOKBodyAgentsItems0 `json:"agents"`
}

// Validate validates this list agents OK body
func (o *ListAgentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsOKBody) validateAgents(formats strfmt.Registry) error {
	if swag.IsZero(o.Agents) { // not required
		return nil
	}

	for i := 0; i < len(o.Agents); i++ {
		if swag.IsZero(o.Agents[i]) { // not required
			continue
		}

		if o.Agents[i] != nil {
			if err := o.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listAgentsOk" + "." + "agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list agents OK body based on the context it is used
func (o *ListAgentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsOKBody) contextValidateAgents(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Agents); i++ {
		if o.Agents[i] != nil {
			if err := o.Agents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listAgentsOk" + "." + "agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBody) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsOKBodyAgentsItems0 list agents OK body agents items0
swagger:model ListAgentsOKBodyAgentsItems0
*/
type ListAgentsOKBodyAgentsItems0 struct {
	// Unique agent identifier.
	AgentID string `json:"agent_id,omitempty"`

	// True if the agent password is set.
	IsAgentPasswordSet bool `json:"is_agent_password_set,omitempty"`

	// Agent type.
	AgentType string `json:"agent_type,omitempty"`

	// AWS Access Key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// True if AWS Secret Key is set.
	IsAWSSecretKeySet bool `json:"is_aws_secret_key_set,omitempty"`

	// Creation timestamp.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Log level for exporter.
	LogLevel string `json:"log_level,omitempty"`

	// Limit query length in QAN.
	MaxQueryLength int32 `json:"max_query_length,omitempty"`

	// Limit query log size in QAN.
	MaxQueryLogSize string `json:"max_query_log_size,omitempty"`

	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `json:"metrics_path,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints.
	MetricsScheme string `json:"metrics_scheme,omitempty"`

	// A unique node identifier.
	NodeID string `json:"node_id,omitempty"`

	// True if password for connecting the agent to the database is set.
	IsPasswordSet bool `json:"is_password_set,omitempty"`

	// The pmm-agent identifier.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// True if query comments parsing is disabled.
	CommentsParsingDisabled bool `json:"comments_parsing_disabled,omitempty"`

	// True if RDS basic metrics are disdabled.
	RDSBasicMetricsDisabled bool `json:"rds_basic_metrics_disabled,omitempty"`

	// True if RDS enhanced metrics are disdabled.
	RDSEnhancedMetricsDisabled bool `json:"rds_enhanced_metrics_disabled,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Actual Agent status.
	Status string `json:"status,omitempty"`

	// Last known table count.
	TableCount int32 `json:"table_count,omitempty"`

	// Tablestats group collectors are disabled if there are more than that number of tables.
	// 0 means tablestats group collectors are always enabled (no limit).
	// Negative value means tablestats group collectors are always disabled.
	TableCountTablestatsGroupLimit int32 `json:"table_count_tablestats_group_limit,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// Last update timestamp.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Agent version.
	Version string `json:"version,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	IsConnected bool `json:"is_connected,omitempty"`

	// azure options
	AzureOptions *ListAgentsOKBodyAgentsItems0AzureOptions `json:"azure_options,omitempty"`

	// mongo db options
	MongoDBOptions *ListAgentsOKBodyAgentsItems0MongoDBOptions `json:"mongo_db_options,omitempty"`

	// mysql options
	MysqlOptions *ListAgentsOKBodyAgentsItems0MysqlOptions `json:"mysql_options,omitempty"`

	// postgresql options
	PostgresqlOptions *ListAgentsOKBodyAgentsItems0PostgresqlOptions `json:"postgresql_options,omitempty"`
}

// Validate validates this list agents OK body agents items0
func (o *ListAgentsOKBodyAgentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAzureOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongoDBOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysqlOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresqlOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) validateAzureOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureOptions) { // not required
		return nil
	}

	if o.AzureOptions != nil {
		if err := o.AzureOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) validateMongoDBOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.MongoDBOptions) { // not required
		return nil
	}

	if o.MongoDBOptions != nil {
		if err := o.MongoDBOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongo_db_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) validateMysqlOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.MysqlOptions) { // not required
		return nil
	}

	if o.MysqlOptions != nil {
		if err := o.MysqlOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mysql_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) validatePostgresqlOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.PostgresqlOptions) { // not required
		return nil
	}

	if o.PostgresqlOptions != nil {
		if err := o.PostgresqlOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postgresql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postgresql_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list agents OK body agents items0 based on the context it is used
func (o *ListAgentsOKBodyAgentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAzureOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMongoDBOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMysqlOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePostgresqlOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) contextValidateAzureOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.AzureOptions != nil {
		if err := o.AzureOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) contextValidateMongoDBOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.MongoDBOptions != nil {
		if err := o.MongoDBOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongo_db_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) contextValidateMysqlOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.MysqlOptions != nil {
		if err := o.MysqlOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mysql_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListAgentsOKBodyAgentsItems0) contextValidatePostgresqlOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.PostgresqlOptions != nil {
		if err := o.PostgresqlOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postgresql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postgresql_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBodyAgentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsOKBodyAgentsItems0AzureOptions list agents OK body agents items0 azure options
swagger:model ListAgentsOKBodyAgentsItems0AzureOptions
*/
type ListAgentsOKBodyAgentsItems0AzureOptions struct {
	// Azure client ID.
	ClientID string `json:"client_id,omitempty"`

	// True if Azure client secret is set.
	IsClientSecretSet bool `json:"is_client_secret_set,omitempty"`

	// Azure resource group.
	ResourceGroup string `json:"resource_group,omitempty"`

	// Azure subscription ID.
	SubscriptionID string `json:"subscription_id,omitempty"`

	// Azure tenant ID.
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this list agents OK body agents items0 azure options
func (o *ListAgentsOKBodyAgentsItems0AzureOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list agents OK body agents items0 azure options based on context it is used
func (o *ListAgentsOKBodyAgentsItems0AzureOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0AzureOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0AzureOptions) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBodyAgentsItems0AzureOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsOKBodyAgentsItems0MongoDBOptions list agents OK body agents items0 mongo DB options
swagger:model ListAgentsOKBodyAgentsItems0MongoDBOptions
*/
type ListAgentsOKBodyAgentsItems0MongoDBOptions struct {
	// True if TLS certificate is set.
	IsTLSCertificateKeySet bool `json:"is_tls_certificate_key_set,omitempty"`

	// True if TLS certificate file password is set.
	IsTLSCertificateKeyFilePasswordSet bool `json:"is_tls_certificate_key_file_password_set,omitempty"`

	// MongoDB auth mechanism.
	AuthenticationMechanism string `json:"authentication_mechanism,omitempty"`

	// MongoDB auth database.
	AuthenticationDatabase string `json:"authentication_database,omitempty"`

	// MongoDB stats collections.
	StatsCollections []string `json:"stats_collections"`

	// MongoDB collections limit.
	CollectionsLimit int32 `json:"collections_limit,omitempty"`

	// True if all collectors are enabled.
	EnableAllCollectors bool `json:"enable_all_collectors,omitempty"`
}

// Validate validates this list agents OK body agents items0 mongo DB options
func (o *ListAgentsOKBodyAgentsItems0MongoDBOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list agents OK body agents items0 mongo DB options based on context it is used
func (o *ListAgentsOKBodyAgentsItems0MongoDBOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0MongoDBOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0MongoDBOptions) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBodyAgentsItems0MongoDBOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsOKBodyAgentsItems0MysqlOptions list agents OK body agents items0 mysql options
swagger:model ListAgentsOKBodyAgentsItems0MysqlOptions
*/
type ListAgentsOKBodyAgentsItems0MysqlOptions struct {
	// True if TLS key is set.
	IsTLSKeySet bool `json:"is_tls_key_set,omitempty"`
}

// Validate validates this list agents OK body agents items0 mysql options
func (o *ListAgentsOKBodyAgentsItems0MysqlOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list agents OK body agents items0 mysql options based on context it is used
func (o *ListAgentsOKBodyAgentsItems0MysqlOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0MysqlOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0MysqlOptions) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBodyAgentsItems0MysqlOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListAgentsOKBodyAgentsItems0PostgresqlOptions list agents OK body agents items0 postgresql options
swagger:model ListAgentsOKBodyAgentsItems0PostgresqlOptions
*/
type ListAgentsOKBodyAgentsItems0PostgresqlOptions struct {
	// True if TLS key is set.
	IsSslKeySet bool `json:"is_ssl_key_set,omitempty"`

	// Limit of databases for auto-discovery.
	AutoDiscoveryLimit int32 `json:"auto_discovery_limit,omitempty"`
}

// Validate validates this list agents OK body agents items0 postgresql options
func (o *ListAgentsOKBodyAgentsItems0PostgresqlOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list agents OK body agents items0 postgresql options based on context it is used
func (o *ListAgentsOKBodyAgentsItems0PostgresqlOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0PostgresqlOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBodyAgentsItems0PostgresqlOptions) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBodyAgentsItems0PostgresqlOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
