package membership

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
	"github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"github.com/syllabix/rollpay/backend/service/membership"

	"github.com/syllabix/rollpay/backend/web/rest"
)

// Controller is responsible for handling organization membership
// requests for the API
type Controller struct {
	srv membership.Service
}

// Register controller handlers to the api
func (ctrl *Controller) Register(api *operation.RollpayAPI) {
	api.OrganizationAddOrgMembersV1Handler = organization.
		AddOrgMembersV1HandlerFunc(ctrl.AddMember)

	api.OrganizationGetOrgMembersV1Handler = organization.
		GetOrgMembersV1HandlerFunc(ctrl.GetMembers)
}

func (ctrl *Controller) GetMembers(params organization.GetOrgMembersV1Params, session *model.Principal) middleware.Responder {
	list, err := ctrl.srv.GetAllByOrgID(params.HTTPRequest.Context(), params.ID)
	switch {
	case err == nil:
		return organization.
			NewGetOrgMembersV1OK().
			WithPayload(&list)

	case errors.Is(err, membership.ErrNotFound):
		return organization.
			NewGetOrgMembersV1NotFound().
			WithPayload(&model.StandardError{
				Message: "hmm... we couldn't quite seem to find what you are looking for",
			})

	default:
		return organization.
			NewGetOrgMembersV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "o wow, something didn't go quite right. please grab a cool beverage and try again",
			})
	}
}

func (ctrl *Controller) AddMember(params organization.AddOrgMembersV1Params, session *model.Principal) middleware.Responder {
	member, err := ctrl.srv.AddMember(params)
	switch {
	case err == nil:
		return organization.
			NewAddOrgMembersV1OK().
			WithPayload(&member)

	case errors.Is(err, membership.ErrInvalid):
		return organization.
			NewAddOrgMembersV1BadRequest().
			WithPayload(&model.StandardError{
				Message: "we were unable to create this membership as either the member or organization are not registered in our systems",
			})

	default:
		return organization.
			NewAddOrgMembersV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "o wow, something didn't go quite right. please grab a cool beverage and try again",
			})
	}
}

// NewController intializes a new api controller for handling health endpoint
// requests
func NewController(srv membership.Service) rest.Controller {
	return rest.MakeController(&Controller{srv})
}
