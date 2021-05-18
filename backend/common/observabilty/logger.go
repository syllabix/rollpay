package observabilty

import (
	"github.com/syllabix/logger"
	"github.com/syllabix/logger/mode"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	logger.Configure(
		logger.AppName("rollpay"),
		logger.Level(zapcore.DebugLevel),
		logger.Mode(mode.Development),
	)

	return logger.New()
}
