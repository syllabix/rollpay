// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/syllabix/rollpay/backend/api/model"
)

// UpdateUserByIDV1HandlerFunc turns a function with the right signature into a update user by ID v1 handler
type UpdateUserByIDV1HandlerFunc func(UpdateUserByIDV1Params, *model.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserByIDV1HandlerFunc) Handle(params UpdateUserByIDV1Params, principal *model.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateUserByIDV1Handler interface for that can handle valid update user by ID v1 params
type UpdateUserByIDV1Handler interface {
	Handle(UpdateUserByIDV1Params, *model.Principal) middleware.Responder
}

// NewUpdateUserByIDV1 creates a new http.Handler for the update user by ID v1 operation
func NewUpdateUserByIDV1(ctx *middleware.Context, handler UpdateUserByIDV1Handler) *UpdateUserByIDV1 {
	return &UpdateUserByIDV1{Context: ctx, Handler: handler}
}

/* UpdateUserByIDV1 swagger:route PUT /v1/user/{id} User updateUserByIdV1

update a user by id

*/
type UpdateUserByIDV1 struct {
	Context *middleware.Context
	Handler UpdateUserByIDV1Handler
}

func (o *UpdateUserByIDV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateUserByIDV1Params()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *model.Principal
	if uprinc != nil {
		principal = uprinc.(*model.Principal) // this is really a model.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
