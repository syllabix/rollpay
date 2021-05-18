package api

import (
	"github.com/syllabix/rollpay/backend/api/controller/authorize"
	"github.com/syllabix/rollpay/backend/api/controller/health"
	"github.com/syllabix/rollpay/backend/api/controller/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewController,
	health.NewController,
	authorize.NewController,
)
