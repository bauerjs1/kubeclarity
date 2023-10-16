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

// NewGetDashboardPackagesPerLanguageParams creates a new GetDashboardPackagesPerLanguageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardPackagesPerLanguageParams() *GetDashboardPackagesPerLanguageParams {
	return &GetDashboardPackagesPerLanguageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardPackagesPerLanguageParamsWithTimeout creates a new GetDashboardPackagesPerLanguageParams object
// with the ability to set a timeout on a request.
func NewGetDashboardPackagesPerLanguageParamsWithTimeout(timeout time.Duration) *GetDashboardPackagesPerLanguageParams {
	return &GetDashboardPackagesPerLanguageParams{
		timeout: timeout,
	}
}

// NewGetDashboardPackagesPerLanguageParamsWithContext creates a new GetDashboardPackagesPerLanguageParams object
// with the ability to set a context for a request.
func NewGetDashboardPackagesPerLanguageParamsWithContext(ctx context.Context) *GetDashboardPackagesPerLanguageParams {
	return &GetDashboardPackagesPerLanguageParams{
		Context: ctx,
	}
}

// NewGetDashboardPackagesPerLanguageParamsWithHTTPClient creates a new GetDashboardPackagesPerLanguageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardPackagesPerLanguageParamsWithHTTPClient(client *http.Client) *GetDashboardPackagesPerLanguageParams {
	return &GetDashboardPackagesPerLanguageParams{
		HTTPClient: client,
	}
}

/*
GetDashboardPackagesPerLanguageParams contains all the parameters to send to the API endpoint

	for the get dashboard packages per language operation.

	Typically these are written to a http.Request.
*/
type GetDashboardPackagesPerLanguageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard packages per language params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardPackagesPerLanguageParams) WithDefaults() *GetDashboardPackagesPerLanguageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard packages per language params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardPackagesPerLanguageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard packages per language params
func (o *GetDashboardPackagesPerLanguageParams) WithTimeout(timeout time.Duration) *GetDashboardPackagesPerLanguageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard packages per language params
func (o *GetDashboardPackagesPerLanguageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard packages per language params
func (o *GetDashboardPackagesPerLanguageParams) WithContext(ctx context.Context) *GetDashboardPackagesPerLanguageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard packages per language params
func (o *GetDashboardPackagesPerLanguageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard packages per language params
func (o *GetDashboardPackagesPerLanguageParams) WithHTTPClient(client *http.Client) *GetDashboardPackagesPerLanguageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard packages per language params
func (o *GetDashboardPackagesPerLanguageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardPackagesPerLanguageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
