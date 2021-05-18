package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/syllabix/rollpay/backend/db"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Store struct {
	db db.Rollpay
}

func (s Store) GetByID(ctx context.Context, id int64) (model.User, error) {
	user, err := model.Users(
		model.UserWhere.ID.EQ(id),
		qm.Load(model.UserRels.LinkedAccounts),
	).One(ctx, s.db)
	if err != nil {
		return failure(err)
	}

	return *user, nil
}

func (s Store) Create(ctx context.Context, user model.User) (model.User, error) {
	err := user.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return failure(err)
	}

	return user, nil
}

func NewStore(db db.Rollpay) Store {
	return Store{db}
}
