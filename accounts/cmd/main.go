package main

import (
	"github.com/transerver/commons/logger"
)

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		logger.Sugar().Fatalf("create accounts app fail: %v", err)
	}

	defer cleanup()
	if err := app.Run(); err != nil {
		logger.Sugar().Fatalf("accounts running error: %v", err)
	}
}
