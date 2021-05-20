package service

import (
	"github.com/syllabix/rollpay/backend/service/membership"
	"github.com/syllabix/rollpay/backend/service/organization"
	"github.com/syllabix/rollpay/backend/service/session"
	"github.com/syllabix/rollpay/backend/service/token"
	"github.com/syllabix/rollpay/backend/service/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewService,
	organization.NewService,
	membership.NewService,
	password.NewManager,
	token.NewService,
	session.NewManager,
	session.NewService,
)
