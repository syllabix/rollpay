package payment

import (
	"github.com/syllabix/rollpay/backend/service/payment/client"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	client.New,
)
