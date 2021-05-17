package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/syllabix/rollpay/backend/db"
)

type Store struct {
	db db.Rollpay
}

func (s Store) GetByID(ctx context.Context, id int64) (model.User, error) {
	user, err := model.FindUser(ctx, s.db, id)
	if err != nil {
		return failure(err)
	}

	return *user, nil
}

func NewStore(db db.Rollpay) Store {
	return Store{db}
}
