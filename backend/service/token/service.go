package token

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/client/payment"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/datastore/user"
)

type Service interface {
	IssueLinkToken(ctx context.Context, userID string) (model.LinkToken, error)
}

type service struct {
	store   user.Store
	payment *payment.Client
}

func (s service) IssueLinkToken(ctx context.Context, userID string) (model.LinkToken, error) {
	uID, err := id.ToInternal(userID)
	if err != nil {
		return failure(err)
	}

	user, err := s.store.GetByID(ctx, uID)
	if err != nil {
		return failure(err)
	}

	token, err := s.payment.CreateLinkToken(ctx, asPaymentUser(user))
	if err != nil {
		return failure(err)
	}

	return asToken(token), nil
}
