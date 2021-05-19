package api

import (
	"github.com/syllabix/rollpay/backend/api/controller/authorize"
	"github.com/syllabix/rollpay/backend/api/controller/health"
	"github.com/syllabix/rollpay/backend/api/controller/membership"
	"github.com/syllabix/rollpay/backend/api/controller/organization"
	"github.com/syllabix/rollpay/backend/api/controller/user"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewController,
	organization.NewController,
	membership.NewController,
	health.NewController,
	authorize.NewController,
)
