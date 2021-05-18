package service

import (
	"github.com/syllabix/rollpay/backend/service/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewService,
	password.NewManager,
)
