package logging

import (
	"context"

	"go.uber.org/zap"
)

// contextKey is a custom type to avoid context key collisions
type contextKey string

var loggerKey contextKey = "_zap_logger"

// WithContext returns the logger associated with the given
// context. If there is no logger, it will return Default.
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Default
	}
	if logger, _ := ctx.Value(loggerKey).(*zap.Logger); logger != nil {
		return logger
	}
	return Default
}

func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
