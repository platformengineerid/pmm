// Code generated by go-swagger; DO NOT EDIT.

package filters

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

// NewGetParams creates a new GetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetParams() *GetParams {
	return &GetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetParamsWithTimeout creates a new GetParams object
// with the ability to set a timeout on a request.
func NewGetParamsWithTimeout(timeout time.Duration) *GetParams {
	return &GetParams{
		timeout: timeout,
	}
}

// NewGetParamsWithContext creates a new GetParams object
// with the ability to set a context for a request.
func NewGetParamsWithContext(ctx context.Context) *GetParams {
	return &GetParams{
		Context: ctx,
	}
}

// NewGetParamsWithHTTPClient creates a new GetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetParamsWithHTTPClient(client *http.Client) *GetParams {
	return &GetParams{
		HTTPClient: client,
	}
}

/*
GetParams contains all the parameters to send to the API endpoint

	for the get operation.

	Typically these are written to a http.Request.
*/
type GetParams struct {
	/* Body.

	   FiltersRequest contains period for which we need filters.
	*/
	Body GetBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParams) WithDefaults() *GetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get params
func (o *GetParams) WithTimeout(timeout time.Duration) *GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get params
func (o *GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get params
func (o *GetParams) WithContext(ctx context.Context) *GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get params
func (o *GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get params
func (o *GetParams) WithHTTPClient(client *http.Client) *GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get params
func (o *GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get params
func (o *GetParams) WithBody(body GetBody) *GetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get params
func (o *GetParams) SetBody(body GetBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
