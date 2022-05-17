package telemetry

import (
	"context"
	"go.uber.org/zap"
)

type loggerKey struct{}

func NewApplicationLogger(ctx context.Context) (*zap.Logger, context.Context) {
	logger, _ := zap.NewProduction()
	return logger, AddLoggerToContext(ctx, logger)
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey{}).(*zap.Logger)
}

func AddLoggerToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
