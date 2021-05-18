package main

import (
	"github.com/syllabix/logger"
	"github.com/syllabix/rollpay/backend"
	"github.com/syllabix/rollpay/backend/common/banner"
	"go.uber.org/zap"
)

func main() {
	banner.Print()

	app := backend.NewApplication()
	if app.Err() != nil {
		logger.New().
			Error("unable to start backend application",
				zap.Error(app.Err()),
			)
	}

	app.Run()
}
