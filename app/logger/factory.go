package logger

import (
	"github.com/charmbracelet/log"
)

var defaultFactory Factory = &factory{}

type Factory interface {
	Default() Logger
	With(keyvals ...any) Logger
	WithPrefix(prefix string) Logger
}

type factory struct{}

func (f *factory) With(keyvals ...any) Logger {
	return log.With(keyvals...)
}

func (f *factory) WithPrefix(prefix string) Logger {
	return log.WithPrefix(prefix)
}

func (f *factory) Default() Logger {
	return log.Default()
}
