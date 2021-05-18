package token

import (
	"context"
	"errors"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/client/payment"
)

type Service struct {
	client payment.Client
}

func IssueLinkToken(ctx context.Context, userID int64) (model.LinkToken, error) {
	return model.LinkToken{}, errors.New("not implemented")
}

func NewService(client payment.Client) Service {
	return Service{client}
}
