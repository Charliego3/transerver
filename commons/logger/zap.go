package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLoggerWithoutWriter retuns logger with io.Discard
func NewLoggerWithoutWriter() (*zap.Logger, func()) {
	return NewLogger(io.Discard)
}

// NewLogger returns logger...
func NewLogger(w io.Writer) (*zap.Logger, func()) {
	var core zapcore.Core
	cfg := zap.NewProductionEncoderConfig()
	cfg.FunctionKey = "F"
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	if w == nil || w == io.Discard {
		cfg.EncodeCaller = zapcore.FullCallerEncoder
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncodeName = zapcore.FullNameEncoder
		cfg.ConsoleSeparator = " "
		encoder := zapcore.NewConsoleEncoder(cfg)
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	} else {
		encoder := zapcore.NewJSONEncoder(cfg)
		core = zapcore.NewCore(encoder, zapcore.AddSync(w), zapcore.InfoLevel)
	}

	logger := zap.New(core)
	return logger, func() {
		_ = logger.Sync()
	}
}
