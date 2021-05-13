package backend

import (
	"time"

	"go.uber.org/fx"
)

type settings struct {
	startTimeout time.Duration
	stopTimeout  time.Duration
}

type Option func(*settings)

// StartTimeout configures the start up timeout of the application
func StartTimeout(dur time.Duration) Option {
	return func(config *settings) {
		config.startTimeout = dur
	}
}

// StopTimeout configures the start up timeout of the application
func StopTimeout(dur time.Duration) Option {
	return func(config *settings) {
		config.stopTimeout = dur
	}
}

func apply(options ...Option) settings {
	config := &settings{
		startTimeout: fx.DefaultTimeout,
		stopTimeout:  fx.DefaultTimeout,
	}

	for _, opt := range options {
		opt(config)
	}

	return *config
}
