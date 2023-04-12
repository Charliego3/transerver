package pkg

import "github.com/charmbracelet/log"

type Option func(*Application)

// WithLoggerOption app using this opts when Logger is not specify
func WithLoggerOption(opt log.Options) Option {
	return func(app *Application) {
		app.loggerOptions = &opt
	}
}

// WithLogger app's Logger will be replace to this logger
func WithLogger(logger *log.Logger) Option {
	return func(app *Application) {
		app.Logger = logger
	}
}
