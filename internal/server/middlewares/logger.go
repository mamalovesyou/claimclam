package middlewares

import (
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

// LoggerMiddleware logs requests with usefull information about the client and how long the request took.
func LoggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Prepare fields to log
			var scheme string
			if r.TLS != nil {
				scheme = "https"
			} else {
				scheme = "http"
			}
			proto := r.Proto
			method := r.Method
			remoteAddr := r.RemoteAddr
			userAgent := r.UserAgent()
			uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

			fields := []zap.Field{
				zap.String("http-scheme", scheme),
				zap.String("http-proto-rest", proto),
				zap.String("http-method", method),
				zap.String("remote-addr", remoteAddr),
				zap.String("user-agent", userAgent),
				zap.String("uri", uri),
			}

			logger.Info("Request started", fields...)

			t1 := time.Now()

			next.ServeHTTP(w, r)

			logger.Info("Request completed",
				append(fields, zap.Float64("elapsed-ms", float64(time.Since(t1).Nanoseconds())/1000000.0))...,
			)
		})
	}
}
