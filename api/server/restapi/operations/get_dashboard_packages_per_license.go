// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDashboardPackagesPerLicenseHandlerFunc turns a function with the right signature into a get dashboard packages per license handler
type GetDashboardPackagesPerLicenseHandlerFunc func(GetDashboardPackagesPerLicenseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDashboardPackagesPerLicenseHandlerFunc) Handle(params GetDashboardPackagesPerLicenseParams) middleware.Responder {
	return fn(params)
}

// GetDashboardPackagesPerLicenseHandler interface for that can handle valid get dashboard packages per license params
type GetDashboardPackagesPerLicenseHandler interface {
	Handle(GetDashboardPackagesPerLicenseParams) middleware.Responder
}

// NewGetDashboardPackagesPerLicense creates a new http.Handler for the get dashboard packages per license operation
func NewGetDashboardPackagesPerLicense(ctx *middleware.Context, handler GetDashboardPackagesPerLicenseHandler) *GetDashboardPackagesPerLicense {
	return &GetDashboardPackagesPerLicense{Context: ctx, Handler: handler}
}

/*
	GetDashboardPackagesPerLicense swagger:route GET /dashboard/packagesPerLicense getDashboardPackagesPerLicense

Get packages count per license type
*/
type GetDashboardPackagesPerLicense struct {
	Context *middleware.Context
	Handler GetDashboardPackagesPerLicenseHandler
}

func (o *GetDashboardPackagesPerLicense) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDashboardPackagesPerLicenseParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
