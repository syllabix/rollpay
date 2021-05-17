package web

import (
	"github.com/syllabix/rollpay/backend/web/docs"
	"github.com/syllabix/rollpay/backend/web/home"
	"github.com/syllabix/rollpay/backend/web/rest"
	"go.uber.org/fx"
)

// Module contains all the dependencies for handling web requests
var Module = fx.Provide(
	home.NewPage,
	NewServer,
	rest.NewServer,
	docs.NewServer,
)
