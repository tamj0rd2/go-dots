package telemetry

import (
	"context"
	"go.uber.org/zap"
)

type loggerKey struct{}

func NewApplicationLogger(ctx context.Context) (*zap.Logger, context.Context) {
	logger, _ := zap.NewProduction()
	return logger, context.WithValue(ctx, loggerKey{}, logger)
}

func GetLogger(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey{}).(*zap.Logger)
}
