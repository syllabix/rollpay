package membership

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	ErrFatal             = errors.New("an unexpected error occurred while accessing the database")
	ErrNotFound          = errors.New("the requested membership(s) do not exist")
	ErrInvalidMembership = errors.New("the requested organization or user does not exist or could be added as a membership")
)

func mapErr(reason error) (err error) {
	var sqlErr *pq.Error
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound

	case errors.As(reason, &sqlErr):
		if sqlErr.Code.Name() == "foreign_key_violation" {
			err = ErrInvalidMembership
			return
		}

		// if it is not an expected SQL error, then use default case
		fallthrough

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return
}

func failure(reason error) (org model.OrganizationMember, err error) {
	return org, mapErr(reason)
}

func rollback(tx boil.Transactor, reason error) (org model.OrganizationMember, err error) {
	err = tx.Rollback()
	if err != nil {
		return org, mapErr(err)
	}

	return org, mapErr(reason)
}
