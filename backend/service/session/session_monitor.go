package session

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
	"go.uber.org/zap"
)

// a monitor wraps and provides observability for a Service
type monitor struct {
	log *zap.Logger
	srv Service
}

func (s monitor) Login(ctx context.Context, email, password string) (model.RollpayToken, error) {
	result, err := s.srv.Login(ctx, email, password)
	if err != nil {
		s.log.Error("login failed", zap.String("reason", err.Error()))
	}
	return result, err
}

func NewService(
	session *Manager,
	password password.Manager,
	users user.Store,
	log *zap.Logger,
) Service {
	return monitor{
		log: log,
		srv: service{session, password, users},
	}
}
