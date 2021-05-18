package user

import (
	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// updatable returns the columns that should be updated
// based upon the state of the user model
func updatable(u model.User) boil.Columns {
	var cols []string

	if len(u.Avatar) > 1 {
		cols = append(cols, model.UserColumns.Avatar)
	}

	if len(u.Email) > 1 {
		cols = append(cols, model.UserColumns.Email)
	}

	if len(u.Password) > 1 {
		cols = append(cols, model.UserColumns.Password)
	}

	if len(u.Username) > 1 {
		cols = append(cols, model.UserColumns.Username)
	}

	return boil.Whitelist(cols...)
}
