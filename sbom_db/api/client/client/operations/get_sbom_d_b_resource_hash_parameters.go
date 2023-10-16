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

// NewGetSbomDBResourceHashParams creates a new GetSbomDBResourceHashParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSbomDBResourceHashParams() *GetSbomDBResourceHashParams {
	return &GetSbomDBResourceHashParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSbomDBResourceHashParamsWithTimeout creates a new GetSbomDBResourceHashParams object
// with the ability to set a timeout on a request.
func NewGetSbomDBResourceHashParamsWithTimeout(timeout time.Duration) *GetSbomDBResourceHashParams {
	return &GetSbomDBResourceHashParams{
		timeout: timeout,
	}
}

// NewGetSbomDBResourceHashParamsWithContext creates a new GetSbomDBResourceHashParams object
// with the ability to set a context for a request.
func NewGetSbomDBResourceHashParamsWithContext(ctx context.Context) *GetSbomDBResourceHashParams {
	return &GetSbomDBResourceHashParams{
		Context: ctx,
	}
}

// NewGetSbomDBResourceHashParamsWithHTTPClient creates a new GetSbomDBResourceHashParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSbomDBResourceHashParamsWithHTTPClient(client *http.Client) *GetSbomDBResourceHashParams {
	return &GetSbomDBResourceHashParams{
		HTTPClient: client,
	}
}

/*
GetSbomDBResourceHashParams contains all the parameters to send to the API endpoint

	for the get sbom d b resource hash operation.

	Typically these are written to a http.Request.
*/
type GetSbomDBResourceHashParams struct {

	// ResourceHash.
	ResourceHash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sbom d b resource hash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSbomDBResourceHashParams) WithDefaults() *GetSbomDBResourceHashParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sbom d b resource hash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSbomDBResourceHashParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) WithTimeout(timeout time.Duration) *GetSbomDBResourceHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) WithContext(ctx context.Context) *GetSbomDBResourceHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) WithHTTPClient(client *http.Client) *GetSbomDBResourceHashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceHash adds the resourceHash to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) WithResourceHash(resourceHash string) *GetSbomDBResourceHashParams {
	o.SetResourceHash(resourceHash)
	return o
}

// SetResourceHash adds the resourceHash to the get sbom d b resource hash params
func (o *GetSbomDBResourceHashParams) SetResourceHash(resourceHash string) {
	o.ResourceHash = resourceHash
}

// WriteToRequest writes these params to a swagger request
func (o *GetSbomDBResourceHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceHash
	if err := r.SetPathParam("resourceHash", o.ResourceHash); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
