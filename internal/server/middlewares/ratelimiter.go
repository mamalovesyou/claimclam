package middlewares

import (
	"net/http"
	"strings"

	"github.com/mamalovesyou/claimclam/internal/ratelimiter"
)

// IPRateLimitMiddleware creates a middleware for HTTP handlers to limit requests
// based on the IP address of the requester.
func IPRateLimitMiddleware(ipLimiter *ratelimiter.IPRateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the limiter for the IP address
			ip := strings.Split(r.RemoteAddr, ":")[0]
			if !ipLimiter.Allow(ip) {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
