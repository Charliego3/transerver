package app

import (
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/logger"
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
