// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/syllabix/rollpay/backend/api/model"
)

// CreateOrganizationV1HandlerFunc turns a function with the right signature into a create organization v1 handler
type CreateOrganizationV1HandlerFunc func(CreateOrganizationV1Params, *model.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOrganizationV1HandlerFunc) Handle(params CreateOrganizationV1Params, principal *model.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateOrganizationV1Handler interface for that can handle valid create organization v1 params
type CreateOrganizationV1Handler interface {
	Handle(CreateOrganizationV1Params, *model.Principal) middleware.Responder
}

// NewCreateOrganizationV1 creates a new http.Handler for the create organization v1 operation
func NewCreateOrganizationV1(ctx *middleware.Context, handler CreateOrganizationV1Handler) *CreateOrganizationV1 {
	return &CreateOrganizationV1{Context: ctx, Handler: handler}
}

/* CreateOrganizationV1 swagger:route POST /v1/organization Organization createOrganizationV1

create a new Organization

*/
type CreateOrganizationV1 struct {
	Context *middleware.Context
	Handler CreateOrganizationV1Handler
}

func (o *CreateOrganizationV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateOrganizationV1Params()
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
