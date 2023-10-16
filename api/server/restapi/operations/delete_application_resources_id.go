// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteApplicationResourcesIDHandlerFunc turns a function with the right signature into a delete application resources ID handler
type DeleteApplicationResourcesIDHandlerFunc func(DeleteApplicationResourcesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteApplicationResourcesIDHandlerFunc) Handle(params DeleteApplicationResourcesIDParams) middleware.Responder {
	return fn(params)
}

// DeleteApplicationResourcesIDHandler interface for that can handle valid delete application resources ID params
type DeleteApplicationResourcesIDHandler interface {
	Handle(DeleteApplicationResourcesIDParams) middleware.Responder
}

// NewDeleteApplicationResourcesID creates a new http.Handler for the delete application resources ID operation
func NewDeleteApplicationResourcesID(ctx *middleware.Context, handler DeleteApplicationResourcesIDHandler) *DeleteApplicationResourcesID {
	return &DeleteApplicationResourcesID{Context: ctx, Handler: handler}
}

/*
	DeleteApplicationResourcesID swagger:route DELETE /applicationResources/{id} deleteApplicationResourcesId

Delete Application Resource
*/
type DeleteApplicationResourcesID struct {
	Context *middleware.Context
	Handler DeleteApplicationResourcesIDHandler
}

func (o *DeleteApplicationResourcesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteApplicationResourcesIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
