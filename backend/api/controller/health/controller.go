package health

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
	"github.com/syllabix/rollpay/backend/api/rest/operation/health"
	"github.com/syllabix/rollpay/backend/web/rest"
)

// Controller is responsible for handling media related request for the API
type Controller struct {
	// generally a controller will have more dependencies
	// they would be added here
}

// Register controller handlers to the api
func (ctrl *Controller) Register(api *operation.RollpayAPI) {
	api.HealthCheckV1Handler = health.CheckV1HandlerFunc(ctrl.HealthCheck)
}

// HealthCheck handles health check requests to the toaster api
func (ctrl *Controller) HealthCheck(params health.CheckV1Params) middleware.Responder {
	// for now we just return a 200
	return health.NewCheckV1OK().
		WithPayload(&model.StandardResponse{
			Message: "ok",
		})
}

// NewController intializes a new api controller for handling health endpoint
// requests
func NewController() rest.Controller {
	return rest.MakeController(&Controller{})
}
