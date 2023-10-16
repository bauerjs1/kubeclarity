// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/openclarity/kubeclarity/sbom_db/api/server/models"
)

// NewPutSbomDBResourceHashParams creates a new PutSbomDBResourceHashParams object
//
// There are no default values defined in the spec.
func NewPutSbomDBResourceHashParams() PutSbomDBResourceHashParams {

	return PutSbomDBResourceHashParams{}
}

// PutSbomDBResourceHashParams contains all the bound params for the put sbom d b resource hash operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutSbomDBResourceHash
type PutSbomDBResourceHashParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.SBOM
	/*
	  Required: true
	  In: path
	*/
	ResourceHash string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutSbomDBResourceHashParams() beforehand.
func (o *PutSbomDBResourceHashParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SBOM
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	rResourceHash, rhkResourceHash, _ := route.Params.GetOK("resourceHash")
	if err := o.bindResourceHash(rResourceHash, rhkResourceHash, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindResourceHash binds and validates parameter ResourceHash from path.
func (o *PutSbomDBResourceHashParams) bindResourceHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ResourceHash = raw

	return nil
}
