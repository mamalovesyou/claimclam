package server

import "go.uber.org/zap"

type Config struct {
	Address        string
	Port           string
	AllowedOrigins []string
	Logger         *zap.Logger
}
