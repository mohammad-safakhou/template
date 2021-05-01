package logger

import (
	"go.uber.org/zap"
)

func ZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}

