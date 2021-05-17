package service

import (
	"github.com/syllabix/rollpay/backend/service/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewService,
)
