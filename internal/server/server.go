package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/fx"
)

func NewMux(lc fx.Lifecycle, cfg *Config) *mux.Router {
	logger := cfg.Logger
	logger.Info("Executing NewMux.")

	r := mux.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS"},
		Debug:            true,
	}).Handler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server.")
			go func() {
				if err := server.ListenAndServe(); err != nil {
					logger.Sugar().Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return r
}
