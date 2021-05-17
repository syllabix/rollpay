package datastore

import (
	"github.com/syllabix/rollpay/backend/datastore/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewStore,
)
