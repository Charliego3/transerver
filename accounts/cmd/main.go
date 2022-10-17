package main

import (
	"github.com/transerver/commons/logger"
)

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		logger.Sugar().Fatal("create accounts app fail", err)
	}

	defer cleanup()
	if err := app.Run(); err != nil {
		logger.Sugar().Fatal("accounts running error", err)
	}
}
