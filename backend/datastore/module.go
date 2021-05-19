package datastore

import (
	"github.com/syllabix/rollpay/backend/datastore/membership"
	"github.com/syllabix/rollpay/backend/datastore/organization"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewStore,
	organization.NewStore,
	membership.NewStore,
)
