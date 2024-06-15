package app

import (
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/pkg/configs"
	"github.com/transerver/pkg/logger"
)

var Name string

func init() {
	if strutil.IsNotBlank(Name) {
		return
	}

	if configs.Bootstrap.Root().Env() != configs.DEV {
		logger.Sugar().Fatal("Is app.Name not injected during build?")
	}

	Name = configs.Bootstrap.Root().Name
}