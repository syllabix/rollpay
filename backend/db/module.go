package db

import "go.uber.org/fx"

// Module contains all of the dependencies needed for direct database
// access, configuration, and migrations
var Module = fx.Provide(
	SetupRollpay,
)
