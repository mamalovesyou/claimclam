package app

import (
	"github.com/gorilla/mux"
	"github.com/mamalovesyou/claimclam/internal/logging"
	"github.com/mamalovesyou/claimclam/internal/podcasts"
	"github.com/mamalovesyou/claimclam/internal/server"
	"github.com/mamalovesyou/claimclam/services/gateway/graph/resolvers"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp(config *server.Config) *fx.App {
	app := fx.New(
		fx.Provide(func(lc fx.Lifecycle) *mux.Router {
			return server.NewMux(lc, config)
		}),
		fx.Provide(
			logging.NewLogger,
			podcasts.NewClient,
			resolvers.NewResolver,
		),
		fx.Invoke(
			resolvers.RegisterRoutes,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	return app
}
