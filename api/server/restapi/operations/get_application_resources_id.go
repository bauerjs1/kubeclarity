// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetApplicationResourcesIDHandlerFunc turns a function with the right signature into a get application resources ID handler
type GetApplicationResourcesIDHandlerFunc func(GetApplicationResourcesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApplicationResourcesIDHandlerFunc) Handle(params GetApplicationResourcesIDParams) middleware.Responder {
	return fn(params)
}

// GetApplicationResourcesIDHandler interface for that can handle valid get application resources ID params
type GetApplicationResourcesIDHandler interface {
	Handle(GetApplicationResourcesIDParams) middleware.Responder
}

// NewGetApplicationResourcesID creates a new http.Handler for the get application resources ID operation
func NewGetApplicationResourcesID(ctx *middleware.Context, handler GetApplicationResourcesIDHandler) *GetApplicationResourcesID {
	return &GetApplicationResourcesID{Context: ctx, Handler: handler}
}

/*
	GetApplicationResourcesID swagger:route GET /applicationResources/{id} getApplicationResourcesId

Get Application Resource
*/
type GetApplicationResourcesID struct {
	Context *middleware.Context
	Handler GetApplicationResourcesIDHandler
}

func (o *GetApplicationResourcesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetApplicationResourcesIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
