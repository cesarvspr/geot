package logger

import (
	"context"

	"github.com/cesarvspr/geot/utils"
	"go.uber.org/zap"
)

type LoggerKey struct{}

var fallbackLogger *zap.SugaredLogger

func init() {
	config := config()
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "severity"
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	if logger, err := config.Build(); err != nil {
		fallbackLogger = zap.NewNop().Sugar()
	} else {
		fallbackLogger = logger.Named("default").Sugar()
	}
	fallbackLogger.Info("Logger started.")
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(LoggerKey{}).(*zap.SugaredLogger); ok {
		return logger
	}
	return fallbackLogger
}

func config() zap.Config {
	if utils.GetEnvironment() == "prod" {
		return zap.NewProductionConfig()
	}
	return zap.NewDevelopmentConfig()
}
