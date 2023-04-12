package logger

import (
	"github.com/transerver/pkg1/configs"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	standard *zap.Logger
	sugar    *zap.SugaredLogger

	standardOnce sync.Once
	sugarOnce    sync.Once
)

func Standard() *zap.Logger {
	standardOnce.Do(func() {
		var core zapcore.Core
		cfg := zap.NewProductionEncoderConfig()
		cfg.FunctionKey = "F"
		cfg.EncodeTime = zapcore.RFC3339TimeEncoder
		cfg.EncodeDuration = zapcore.StringDurationEncoder
		if configs.Bootstrap.Root().Env() == configs.DEV {
			cfg.EncodeCaller = zapcore.FullCallerEncoder
			cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
			cfg.EncodeName = zapcore.FullNameEncoder
			cfg.ConsoleSeparator = " "
			encoder := zapcore.NewConsoleEncoder(cfg)
			core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
		} else {
			encoder := zapcore.NewJSONEncoder(cfg)
			core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
		}

		standard = zap.New(core)
	})
	return standard
}

func Sugar() *zap.SugaredLogger {
	sugarOnce.Do(func() {
		sugar = Standard().Sugar()
	})
	return sugar
}
