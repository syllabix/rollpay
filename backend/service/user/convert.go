package user

import (
	"strconv"
	"strings"

	"github.com/go-openapi/strfmt"
	api "github.com/syllabix/rollpay/backend/api/model"
	db "github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/volatiletech/null/v8"
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

func asNewUser(u api.User, password []byte) db.User {
	return db.User{
		Email:    strings.ToLower(u.Email.String()),
		Password: string(password),
		Username: u.Username,
		Avatar:   null.BytesFrom([]byte(u.Avatar.String())),
	}
}
