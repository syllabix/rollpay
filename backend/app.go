package backend

import (
	"context"

	"github.com/syllabix/rollpay/backend/config"
	"github.com/syllabix/rollpay/backend/util/observabilty"
	"github.com/syllabix/rollpay/backend/util/profiler"
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
		fx.StartTimeout(settings.startTimeout),
		fx.StopTimeout(settings.stopTimeout),
		// the default fx logger will print the dependency
		// graph for debugging purposes. it creates quite
		// a bit of noise on start up so it is disabled for now
		fx.NopLogger,

		// application dependencies
		fx.Provide(config.Load),
		web.Module,
		observabilty.Module,

		fx.Invoke(start),

		// enables pprof on default port 6060
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
