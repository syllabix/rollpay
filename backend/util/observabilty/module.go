package observabilty

import "go.uber.org/fx"

// Module exposes all dependencies that facilitate obervability
// (logging, metrics, etc)
var Module = fx.Provide(
	NewLogger,
)
