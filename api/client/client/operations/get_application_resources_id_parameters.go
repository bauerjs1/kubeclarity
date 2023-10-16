// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetApplicationResourcesIDParams creates a new GetApplicationResourcesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetApplicationResourcesIDParams() *GetApplicationResourcesIDParams {
	return &GetApplicationResourcesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationResourcesIDParamsWithTimeout creates a new GetApplicationResourcesIDParams object
// with the ability to set a timeout on a request.
func NewGetApplicationResourcesIDParamsWithTimeout(timeout time.Duration) *GetApplicationResourcesIDParams {
	return &GetApplicationResourcesIDParams{
		timeout: timeout,
	}
}

// NewGetApplicationResourcesIDParamsWithContext creates a new GetApplicationResourcesIDParams object
// with the ability to set a context for a request.
func NewGetApplicationResourcesIDParamsWithContext(ctx context.Context) *GetApplicationResourcesIDParams {
	return &GetApplicationResourcesIDParams{
		Context: ctx,
	}
}

// NewGetApplicationResourcesIDParamsWithHTTPClient creates a new GetApplicationResourcesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetApplicationResourcesIDParamsWithHTTPClient(client *http.Client) *GetApplicationResourcesIDParams {
	return &GetApplicationResourcesIDParams{
		HTTPClient: client,
	}
}

/*
GetApplicationResourcesIDParams contains all the parameters to send to the API endpoint

	for the get application resources ID operation.

	Typically these are written to a http.Request.
*/
type GetApplicationResourcesIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get application resources ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationResourcesIDParams) WithDefaults() *GetApplicationResourcesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get application resources ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationResourcesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get application resources ID params
func (o *GetApplicationResourcesIDParams) WithTimeout(timeout time.Duration) *GetApplicationResourcesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application resources ID params
func (o *GetApplicationResourcesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application resources ID params
func (o *GetApplicationResourcesIDParams) WithContext(ctx context.Context) *GetApplicationResourcesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application resources ID params
func (o *GetApplicationResourcesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application resources ID params
func (o *GetApplicationResourcesIDParams) WithHTTPClient(client *http.Client) *GetApplicationResourcesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application resources ID params
func (o *GetApplicationResourcesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get application resources ID params
func (o *GetApplicationResourcesIDParams) WithID(id string) *GetApplicationResourcesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get application resources ID params
func (o *GetApplicationResourcesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationResourcesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
