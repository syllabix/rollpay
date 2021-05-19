package membership

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

func (s Store) GetAllByOrgID(ctx context.Context, orgID int64) (model.OrganizationMemberSlice, error) {
	members, err := model.OrganizationMembers(
		model.OrganizationMemberWhere.OrganizationID.EQ(orgID),
		qm.Load(model.OrganizationMemberRels.User),
		qm.Load(model.OrganizationMemberRels.Organization),
	).All(ctx, s.db)
	if err != nil {
		return nil, mapErr(err)
	}

	return members, nil
}

func (s Store) AddMember(ctx context.Context, member model.OrganizationMember) (model.OrganizationMember, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return failure(err)
	}

	var (
		conflictCols = []string{
			model.OrganizationMemberColumns.UserID,
			model.OrganizationMemberColumns.OrganizationID,
		}
	)

	err = member.Upsert(ctx, tx, true, conflictCols, boil.Infer(), boil.Infer())
	if err != nil {
		return rollback(tx, err)
	}

	updated, err := model.OrganizationMembers(
		model.OrganizationMemberWhere.UserID.EQ(member.UserID),
		model.OrganizationMemberWhere.OrganizationID.EQ(member.OrganizationID),
		qm.Load(model.OrganizationMemberRels.User),
		qm.Load(model.OrganizationMemberRels.Organization),
	).One(ctx, tx)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollback(tx, err)
	}

	return *updated, nil
}

func NewStore(db db.Rollpay) Store {
	return Store{db}
}
