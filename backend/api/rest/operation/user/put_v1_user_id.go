// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutV1UserIDHandlerFunc turns a function with the right signature into a put v1 user ID handler
type PutV1UserIDHandlerFunc func(PutV1UserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutV1UserIDHandlerFunc) Handle(params PutV1UserIDParams) middleware.Responder {
	return fn(params)
}

// PutV1UserIDHandler interface for that can handle valid put v1 user ID params
type PutV1UserIDHandler interface {
	Handle(PutV1UserIDParams) middleware.Responder
}

// NewPutV1UserID creates a new http.Handler for the put v1 user ID operation
func NewPutV1UserID(ctx *middleware.Context, handler PutV1UserIDHandler) *PutV1UserID {
	return &PutV1UserID{Context: ctx, Handler: handler}
}

/* PutV1UserID swagger:route PUT /v1/user/{id} User putV1UserId

update a user by id

*/
type PutV1UserID struct {
	Context *middleware.Context
	Handler PutV1UserIDHandler
}

func (o *PutV1UserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutV1UserIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
