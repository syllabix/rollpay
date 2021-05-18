package client

import (
	"github.com/syllabix/rollpay/backend/common/client/payment"
	"go.uber.org/fx"
)

// Module exposes all third party clients
// for injection
var Module = fx.Provide(
	payment.NewClient,
)
