package token

import (
	"github.com/go-openapi/strfmt"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/client/payment"
	"github.com/syllabix/rollpay/backend/common/id"
	db "github.com/syllabix/rollpay/backend/datastore/model"
)

func asPaymentUser(u db.User) payment.User {
	return payment.User{
		ID:       id.AsPublic(u.ID),
		Email:    u.Email,
		Language: "en", // tears...
	}
}

func asToken(tk payment.LinkToken) model.LinkToken {
	return model.LinkToken{
		Token:      tk.Token,
		Expiration: strfmt.DateTime(tk.Expiration),
	}
}
