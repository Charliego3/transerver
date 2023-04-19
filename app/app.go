package app

import (
	"github.com/transerver/app/logger"
)

type Application struct {
	servers []Server
}

// NewApp returns Application
func NewApp(opts ...Option) *Application {
	app := &Application{}
	app.getOpts(opts...)
	return app
}

// getOpts handling and aggregation options
func (app *Application) getOpts(opts ...Option) {
	for _, opt := range opts {
		opt(app)
	}
}

func (app *Application) Start() chan struct{} {
	return make(chan struct{})
}

func (app *Application) Run() {
	logger.Debugf("application is running...  %s", "replacement")
}
