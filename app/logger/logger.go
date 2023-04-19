package logger

import (
	"fmt"
	"sync"
)

type Logger interface {
	Debug(msg interface{}, keyvals ...interface{})
	Info(msg interface{}, keyvals ...interface{})
	Warn(msg interface{}, keyvals ...interface{})
	Error(msg interface{}, keyvals ...interface{})
	Fatal(msg interface{}, keyvals ...interface{})
	Print(msg interface{}, keyvals ...interface{})
}

type Loggerf interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

var (
	defaultLogger Logger
	loggerOnce    sync.Once
)

func With(keyvals ...any) Logger {
	return defaultFactory.With(keyvals...)
}

func WithPrefix(prefix string) Logger {
	return defaultFactory.WithPrefix(prefix)
}

func SetFactory(factory Factory) {
	defaultFactory = factory
}

func SetLogger(logger Logger) {
	defaultLogger = logger
}

func getLogger() Logger {
	loggerOnce.Do(func() {
		defaultLogger = defaultFactory.Default()
	})
	return defaultLogger
}

func Debug(msg interface{}, keyvals ...interface{}) {
	getLogger().Debug(msg, keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	getLogger().Info(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	getLogger().Warn(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	getLogger().Error(msg, keyvals...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	getLogger().Fatal(msg, keyvals...)
}

func Print(msg interface{}, keyvals ...interface{}) {
	getLogger().Print(msg, keyvals...)
}

func Debugf(format string, args ...interface{}) {
	logger := getLogger()
	if l, ok := logger.(interface {
		Debugf(string, ...any)
	}); ok {
		l.Debugf(format, args...)
	} else {
		Debug(fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	logger := getLogger()
	if l, ok := logger.(interface {
		Infof(string, ...any)
	}); ok {
		l.Infof(format, args...)
	} else {
		Info(fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	logger := getLogger()
	if l, ok := logger.(interface {
		Warnf(string, ...any)
	}); ok {
		l.Warnf(format, args...)
	} else {
		Warn(fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	logger := getLogger()
	if l, ok := logger.(interface {
		Errorf(string, ...any)
	}); ok {
		l.Errorf(format, args...)
	} else {
		Error(fmt.Sprintf(format, args...))
	}
}

func Fatalf(format string, args ...interface{}) {
	logger := getLogger()
	if l, ok := logger.(interface {
		Fatalf(string, ...any)
	}); ok {
		l.Fatalf(format, args...)
	} else {
		Fatal(fmt.Sprintf(format, args...))
	}
}

func Printf(format string, args ...interface{}) {
	logger := getLogger()
	if l, ok := logger.(interface {
		Printf(string, ...any)
	}); ok {
		l.Printf(format, args...)
	} else {
		Print(fmt.Sprintf(format, args...))
	}
}
