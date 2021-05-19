package organization

import (
	"context"
	"database/sql"

	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/syllabix/rollpay/backend/db"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Store struct {
	db db.Rollpay
}

func (s Store) GetByID(ctx context.Context, id int64) (model.Organization, error) {
	org, err := model.Organizations(
		model.OrganizationWhere.ID.EQ(id),
		qm.Load(model.OrganizationRels.OrganizationAccounts,
			qm.Load(model.OrganizationAccountRels.LinkedAccount)),
	).One(ctx, s.db)
	if err != nil {
		return failure(err)
	}

	return *org, nil
}

func (s Store) Create(ctx context.Context, org model.Organization) (model.Organization, error) {
	err := org.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return failure(err)
	}

	return org, nil
}

func (s Store) Update(ctx context.Context, org model.Organization) (model.Organization, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return failure(err)
	}

	effected, err := org.Update(ctx, tx, updatable(org))
	if err != nil {
		return rollback(tx, err)
	}

	if effected < 1 {
		return rollback(tx, sql.ErrNoRows)
	}

	err = org.Reload(ctx, tx)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollback(tx, err)
	}

	return org, nil
}

func (s Store) Delete(ctx context.Context, id int64) error {
	effected, err := model.Organizations(
		model.OrganizationWhere.ID.EQ(id),
	).DeleteAll(ctx, s.db, false)

	if err != nil {
		return ErrFatal
	}

	if effected < 1 {
		return ErrNotFound
	}

	return nil
}

func NewStore(db db.Rollpay) Store {
	return Store{db}
}
