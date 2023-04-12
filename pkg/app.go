package pkg

import (
	"github.com/charmbracelet/log"
	"os"
	"time"
)

type Application struct {
	loggerOptions *log.Options
	Logger        *log.Logger
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

	app.Logger = DefaultFunc(app.Logger, func() *log.Logger {
		app.loggerOptions = Default(app.loggerOptions, &log.Options{
			TimeFormat:      time.DateTime,
			Level:           log.InfoLevel,
			Prefix:          "Application",
			ReportTimestamp: true,
			ReportCaller:    true,
		})
		return log.NewWithOptions(os.Stdout, *app.loggerOptions)
	})
}

func (app *Application) Run() {
	app.Logger.Info("application is running...")
}
