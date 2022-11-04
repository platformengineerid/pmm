// Code generated by go-swagger; DO NOT EDIT.

package security_checks

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

// NewListFailedServicesParams creates a new ListFailedServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListFailedServicesParams() *ListFailedServicesParams {
	return &ListFailedServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListFailedServicesParamsWithTimeout creates a new ListFailedServicesParams object
// with the ability to set a timeout on a request.
func NewListFailedServicesParamsWithTimeout(timeout time.Duration) *ListFailedServicesParams {
	return &ListFailedServicesParams{
		timeout: timeout,
	}
}

// NewListFailedServicesParamsWithContext creates a new ListFailedServicesParams object
// with the ability to set a context for a request.
func NewListFailedServicesParamsWithContext(ctx context.Context) *ListFailedServicesParams {
	return &ListFailedServicesParams{
		Context: ctx,
	}
}

// NewListFailedServicesParamsWithHTTPClient creates a new ListFailedServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListFailedServicesParamsWithHTTPClient(client *http.Client) *ListFailedServicesParams {
	return &ListFailedServicesParams{
		HTTPClient: client,
	}
}

/*
ListFailedServicesParams contains all the parameters to send to the API endpoint

	for the list failed services operation.

	Typically these are written to a http.Request.
*/
type ListFailedServicesParams struct {
	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list failed services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFailedServicesParams) WithDefaults() *ListFailedServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list failed services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFailedServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list failed services params
func (o *ListFailedServicesParams) WithTimeout(timeout time.Duration) *ListFailedServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list failed services params
func (o *ListFailedServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list failed services params
func (o *ListFailedServicesParams) WithContext(ctx context.Context) *ListFailedServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list failed services params
func (o *ListFailedServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list failed services params
func (o *ListFailedServicesParams) WithHTTPClient(client *http.Client) *ListFailedServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list failed services params
func (o *ListFailedServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list failed services params
func (o *ListFailedServicesParams) WithBody(body interface{}) *ListFailedServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list failed services params
func (o *ListFailedServicesParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListFailedServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
