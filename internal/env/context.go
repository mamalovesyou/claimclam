package env

import (
	"context"
)

type contextKey string

var environmentKey contextKey = "_environment"

func FromContext(ctx context.Context) *string {
	if ctx == nil {
		return nil
	}
	if env, _ := ctx.Value(environmentKey).(*string); env != nil {
		return env
	}
	return nil
}

func ContextWithEnv(ctx context.Context, env string) context.Context {
	return context.WithValue(ctx, environmentKey, env)
}
