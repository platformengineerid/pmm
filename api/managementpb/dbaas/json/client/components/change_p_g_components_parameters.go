// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewChangePGComponentsParams creates a new ChangePGComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangePGComponentsParams() *ChangePGComponentsParams {
	return &ChangePGComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangePGComponentsParamsWithTimeout creates a new ChangePGComponentsParams object
// with the ability to set a timeout on a request.
func NewChangePGComponentsParamsWithTimeout(timeout time.Duration) *ChangePGComponentsParams {
	return &ChangePGComponentsParams{
		timeout: timeout,
	}
}

// NewChangePGComponentsParamsWithContext creates a new ChangePGComponentsParams object
// with the ability to set a context for a request.
func NewChangePGComponentsParamsWithContext(ctx context.Context) *ChangePGComponentsParams {
	return &ChangePGComponentsParams{
		Context: ctx,
	}
}

// NewChangePGComponentsParamsWithHTTPClient creates a new ChangePGComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangePGComponentsParamsWithHTTPClient(client *http.Client) *ChangePGComponentsParams {
	return &ChangePGComponentsParams{
		HTTPClient: client,
	}
}

/*
ChangePGComponentsParams contains all the parameters to send to the API endpoint

	for the change p g components operation.

	Typically these are written to a http.Request.
*/
type ChangePGComponentsParams struct {
	// Body.
	Body ChangePGComponentsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change p g components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePGComponentsParams) WithDefaults() *ChangePGComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change p g components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePGComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change p g components params
func (o *ChangePGComponentsParams) WithTimeout(timeout time.Duration) *ChangePGComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change p g components params
func (o *ChangePGComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change p g components params
func (o *ChangePGComponentsParams) WithContext(ctx context.Context) *ChangePGComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change p g components params
func (o *ChangePGComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change p g components params
func (o *ChangePGComponentsParams) WithHTTPClient(client *http.Client) *ChangePGComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change p g components params
func (o *ChangePGComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change p g components params
func (o *ChangePGComponentsParams) WithBody(body ChangePGComponentsBody) *ChangePGComponentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change p g components params
func (o *ChangePGComponentsParams) SetBody(body ChangePGComponentsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePGComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}