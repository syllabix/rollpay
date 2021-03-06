package backend

import (
	"context"

	"github.com/syllabix/rollpay/backend/api"
	"github.com/syllabix/rollpay/backend/common/client"
	"github.com/syllabix/rollpay/backend/common/observabilty"
	"github.com/syllabix/rollpay/backend/common/profiler"
	"github.com/syllabix/rollpay/backend/config"
	"github.com/syllabix/rollpay/backend/datastore"
	"github.com/syllabix/rollpay/backend/db"
	"github.com/syllabix/rollpay/backend/service"
	"github.com/syllabix/rollpay/backend/service/payment"
	"github.com/syllabix/rollpay/backend/web"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Application is a web service that powers the
// rollpay backend. It handles application start,
// dependency injection, and graceful shutdowns
// It wraps an uber fx app, so for detailed documentation
// please check the go doc at https://pkg.go.dev/go.uber.org/fx
type Application struct {
	*fx.App
}

// NewApplication constucts a rollpay backend
// application ready to be run.
func NewApplication(options ...Option) Application {
	settings := apply(options...)

	app := fx.New(
		// fx configuration
		fx.StartTimeout(settings.startTimeout),
		fx.StopTimeout(settings.stopTimeout),
		// the default fx logger will print the dependency
		// graph for debugging purposes. it creates quite
		// a bit of noise on start up so it is disabled for now
		fx.NopLogger,

		// application dependencies
		fx.Provide(config.Load),
		db.Module,
		api.Module,
		web.Module,
		observabilty.Module,
		payment.Module,
		service.Module,
		datastore.Module,
		client.Module,

		// start the engines
		fx.Invoke(start),
		fx.Invoke(profiler.Start),
	)

	return Application{app}
}

// start boots up the web server hooking it into application start and stop routines
// enabling graceful shutdown
func start(lc fx.Lifecycle, server web.Server, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			log.Info("web server listening and ready to accept connections",
				zap.String("address", server.Addr))
			return nil
		},

		OnStop: func(ctx context.Context) error {
			log.Info("powering down rollpay web service...")
			return server.Shutdown(ctx)
		},
	})
}
