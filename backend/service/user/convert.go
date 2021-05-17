package user

import (
	"strconv"

	"github.com/go-openapi/strfmt"
	api "github.com/syllabix/rollpay/backend/api/model"
	db "github.com/syllabix/rollpay/backend/datastore/model"
)

// asModel converts a db model to an api model
func asModel(u db.User) api.User {
	return api.User{
		ID:        strconv.FormatInt(u.ID, 10),
		Email:     strfmt.Email(u.Email),
		Username:  u.Username,
		Avatar:    u.Avatar.Bytes,
		CreatedAt: strfmt.DateTime(u.CreatedAt),
		UpdatedAt: strfmt.DateTime(u.UpdatedAt),
	}
}
