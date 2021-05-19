package organization

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	ErrNotFound  = errors.New("the requested organization was not found")
	ErrFatal     = errors.New("an unexpected error occurred while accessing the database")
	ErrNameTaken = errors.New("the organization name is already in use")
)

func mapErr(reason error) (err error) {
	var sqlErr *pq.Error
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound

	case errors.As(reason, &sqlErr):
		if sqlErr.Code.Name() == "unique_violation" {
			err = ErrNameTaken
			return
		}

		// if it is not an expected SQL error, then use default case
		fallthrough

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return
}

func failure(reason error) (org model.Organization, err error) {
	return org, mapErr(reason)
}

func rollback(tx boil.Transactor, reason error) (org model.Organization, err error) {
	err = tx.Rollback()
	if err != nil {
		return org, mapErr(err)
	}

	return org, mapErr(reason)
}
