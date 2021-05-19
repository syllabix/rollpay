package organization

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
	"github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	org "github.com/syllabix/rollpay/backend/service/organization"
	"github.com/syllabix/rollpay/backend/web/rest"
)

// Controller is responsible for handling organization related request for the API
type Controller struct {
	srv org.Service
}

// Register controller handlers to the api
func (ctrl *Controller) Register(api *operation.RollpayAPI) {
	api.OrganizationCreateOrganizationV1Handler = organization.
		CreateOrganizationV1HandlerFunc(ctrl.Create)

	api.OrganizationGetOrganizationByIDV1Handler = organization.
		GetOrganizationByIDV1HandlerFunc(ctrl.GetByID)

	api.OrganizationGetAllOrgsV1Handler = organization.
		GetAllOrgsV1HandlerFunc(ctrl.GetAll)
}

func (ctrl *Controller) GetByID(params organization.GetOrganizationByIDV1Params, session *model.Principal) middleware.Responder {
	result, err := ctrl.srv.Get(params.HTTPRequest.Context(), params.ID)
	switch {
	case err == nil:
		return organization.NewGetOrganizationByIDV1OK().
			WithPayload(&result)

	case errors.Is(err, org.ErrNotFound):
		return organization.NewGetOrganizationByIDV1NotFound().
			WithPayload(&model.StandardError{
				Message: "hmm... we were unable to find the organization you were looking for",
			})

	case errors.Is(err, org.ErrInvalid):
		return organization.NewGetOrganizationByIDV1BadRequest().
			WithPayload(&model.StandardError{
				Message: "we are unable to process your request",
			})

	default:
		return organization.NewGetOrganizationByIDV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "o wow, something didn't go quite right. please grab a cool beverage and try again",
			})
	}
}

func (ctrl *Controller) GetAll(params organization.GetAllOrgsV1Params, session *model.Principal) middleware.Responder {
	results, err := ctrl.srv.GetAll(params.HTTPRequest.Context())
	switch {
	case err == nil:
		return organization.NewGetAllOrgsV1OK().
			WithPayload(&model.OrganizationList{
				Results: results,
			})

	default:
		return organization.NewGetAllOrgsV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "oops, something didn't work out as expected on our end. give our systems a minute or two and try again",
			})
	}
}

func (ctrl *Controller) Create(params organization.CreateOrganizationV1Params, session *model.Principal) middleware.Responder {
	newOrg, err := ctrl.srv.Create(params)
	switch {
	case err == nil:
		return organization.NewCreateOrganizationV1Created().
			WithPayload(&newOrg)

	case errors.Is(err, org.ErrNameReserved):
		return organization.NewCreateOrganizationV1Conflict().
			WithPayload(&model.StandardError{
				Message: "sorry, but it looks like this name is already in use.",
			})

	default:
		return organization.NewCreateOrganizationV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "oops, something didn't work out as expected on our end. give our systems a minute or two and try again",
			})
	}
}

// NewController intializes a new api controller for handling health endpoint
// requests
func NewController(srv org.Service) rest.Controller {
	return rest.MakeController(&Controller{srv})
}
