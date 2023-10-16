// Code generated by go-swagger; DO NOT EDIT.

package component

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

// NewDeleteComponentParams creates a new DeleteComponentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteComponentParams() *DeleteComponentParams {
	return &DeleteComponentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteComponentParamsWithTimeout creates a new DeleteComponentParams object
// with the ability to set a timeout on a request.
func NewDeleteComponentParamsWithTimeout(timeout time.Duration) *DeleteComponentParams {
	return &DeleteComponentParams{
		timeout: timeout,
	}
}

// NewDeleteComponentParamsWithContext creates a new DeleteComponentParams object
// with the ability to set a context for a request.
func NewDeleteComponentParamsWithContext(ctx context.Context) *DeleteComponentParams {
	return &DeleteComponentParams{
		Context: ctx,
	}
}

// NewDeleteComponentParamsWithHTTPClient creates a new DeleteComponentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteComponentParamsWithHTTPClient(client *http.Client) *DeleteComponentParams {
	return &DeleteComponentParams{
		HTTPClient: client,
	}
}

/*
DeleteComponentParams contains all the parameters to send to the API endpoint

	for the delete component operation.

	Typically these are written to a http.Request.
*/
type DeleteComponentParams struct {

	/* UUID.

	   The UUID of the component to delete
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteComponentParams) WithDefaults() *DeleteComponentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteComponentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete component params
func (o *DeleteComponentParams) WithTimeout(timeout time.Duration) *DeleteComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete component params
func (o *DeleteComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete component params
func (o *DeleteComponentParams) WithContext(ctx context.Context) *DeleteComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete component params
func (o *DeleteComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete component params
func (o *DeleteComponentParams) WithHTTPClient(client *http.Client) *DeleteComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete component params
func (o *DeleteComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete component params
func (o *DeleteComponentParams) WithUUID(uuid string) *DeleteComponentParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete component params
func (o *DeleteComponentParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
