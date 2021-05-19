package user

import (
	"github.com/go-openapi/strfmt"
	api "github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation/user"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/common/media"
	db "github.com/syllabix/rollpay/backend/datastore/model"
)

// asModel converts a db model to an api model
func asModel(u db.User) api.User {
	return api.User{
		ID:        id.AsPublic(u.ID),
		Email:     strfmt.Email(u.Email),
		Username:  u.Username,
		Avatar:    u.Avatar,
		Accounts:  mapAccounts(u),
		CreatedAt: strfmt.DateTime(u.CreatedAt),
		UpdatedAt: strfmt.DateTime(u.UpdatedAt),
	}
}

func mapAccounts(usr db.User) []*api.LinkedAccount {
	if usr.R == nil {
		return []*api.LinkedAccount{}
	}

	accounts := make([]*api.LinkedAccount, len(usr.R.UserAccounts))
	for i, account := range usr.R.UserAccounts {
		link := account.R.LinkedAccount
		accounts[i] = &api.LinkedAccount{
			ID:        id.AsPublic(link.ID),
			Alias:     link.Alias,
			CreatedAt: strfmt.DateTime(link.CreatedAt),
			UpdatedAt: strfmt.DateTime(link.UpdatedAt),
		}
	}
	return accounts
}

func asNewUser(params user.CreateUserV1Params, password []byte) (db.User, error) {
	img, err := media.Process(params.HTTPRequest.Context(), params.Avatar)
	if err != nil {
		return db.User{}, err
	}

	return db.User{
		Username: params.Username,
		Email:    params.Email.String(),
		Password: string(password),
		Avatar:   img,
	}, nil
}
