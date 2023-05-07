package logger

import (
	"fmt"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type Logger interface {
	Debug(msg interface{}, keyvals ...interface{})
	Info(msg interface{}, keyvals ...interface{})
	Warn(msg interface{}, keyvals ...interface{})
	Error(msg interface{}, keyvals ...interface{})
	Fatal(msg interface{}, keyvals ...interface{})
	Print(msg interface{}, keyvals ...interface{})
	GetLevel() Level
	SetLevel(Level)
	SetTimeFormat(string)
}

type Loggerf interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

type Level = log.Level

var (
	DebugLevel Level = log.DebugLevel
	InfoLevel  Level = log.InfoLevel
	WarnLevel  Level = log.WarnLevel
	ErrorLevel Level = log.ErrorLevel
	FatalLevel Level = log.FatalLevel
)

var (
	defaultLogger Logger
	loggerOnce    sync.Once
)

func SetLogger(logger Logger) {
	defaultLogger = logger
}

func With(keyvals ...any) Logger {
	return defaultFactory.With(keyvals...)
}

func WithPrefix(prefix string) Logger {
	return defaultFactory.WithPrefix(prefix)
}

func Default() Logger {
	loggerOnce.Do(func() {
		defaultLogger = defaultFactory.Default()
		defaultLogger.SetTimeFormat(time.DateTime)
	})
	return defaultLogger
}

func GetLevel() Level {
	return Default().GetLevel()
}

func SetLevel(l Level) {
	Default().SetLevel(l)
}

func SetTimeFormatter(format string) {
	Default().SetTimeFormat(format)
}

func Debug(msg interface{}, keyvals ...interface{}) {
	Default().Debug(msg, keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	Default().Info(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	Default().Warn(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	Default().Error(msg, keyvals...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	Default().Fatal(msg, keyvals...)
}

func Print(msg interface{}, keyvals ...interface{}) {
	Default().Print(msg, keyvals...)
}

func Debugf(format string, args ...interface{}) {
	logger := Default()
	if l, ok := logger.(interface {
		Debugf(string, ...any)
	}); ok {
		l.Debugf(format, args...)
	} else {
		Debug(fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	logger := Default()
	if l, ok := logger.(interface {
		Infof(string, ...any)
	}); ok {
		l.Infof(format, args...)
	} else {
		Info(fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	logger := Default()
	if l, ok := logger.(interface {
		Warnf(string, ...any)
	}); ok {
		l.Warnf(format, args...)
	} else {
		Warn(fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	logger := Default()
	if l, ok := logger.(interface {
		Errorf(string, ...any)
	}); ok {
		l.Errorf(format, args...)
	} else {
		Error(fmt.Sprintf(format, args...))
	}
}

func Fatalf(format string, args ...interface{}) {
	logger := Default()
	if l, ok := logger.(interface {
		Fatalf(string, ...any)
	}); ok {
		l.Fatalf(format, args...)
	} else {
		Fatal(fmt.Sprintf(format, args...))
	}
}

func Printf(format string, args ...interface{}) {
	logger := Default()
	if l, ok := logger.(interface {
		Printf(string, ...any)
	}); ok {
		l.Printf(format, args...)
	} else {
		Print(fmt.Sprintf(format, args...))
	}
}
