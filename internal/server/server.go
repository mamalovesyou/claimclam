package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mamalovesyou/claimclam/internal/ratelimiter"
	"github.com/mamalovesyou/claimclam/internal/server/middlewares"
	"github.com/rs/cors"
	"go.uber.org/fx"
)

func NewMux(lc fx.Lifecycle, cfg *Config) *mux.Router {
	logger := cfg.Logger
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Add rate limiter
	limiter := ratelimiter.NewIPRateLimiter(3, 5)

	// Middlewares
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS"},
	}).Handler)
	r.Use(middlewares.IPRateLimitMiddleware(limiter))
	r.Use(middlewares.LoggerMiddleware(logger))

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
