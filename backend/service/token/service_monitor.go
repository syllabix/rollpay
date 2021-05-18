package token

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/client/payment"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"go.uber.org/zap"
)

type monitor struct {
	log *zap.Logger

	srv Service
}

func (s monitor) IssueLinkToken(ctx context.Context, userID string) (model.LinkToken, error) {
	result, err := s.srv.IssueLinkToken(ctx, userID)
	if err != nil {
		s.log.Error("failed to create link token", zap.Error(err))
	}

	return result, err
}

func NewService(store user.Store, client *payment.Client, log *zap.Logger) Service {
	return monitor{
		log: log,
		srv: service{store, client},
	}
}
