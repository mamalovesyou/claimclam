package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mamalovesyou/claimclam/internal/ratelimiter"
	"github.com/mamalovesyou/claimclam/internal/server/middlewares"
)

func TestIPRateLimitMiddleware(t *testing.T) {
	ipLimiter := ratelimiter.NewIPRateLimiter(1, 1) // Set low limit for testing
	middleware := middlewares.IPRateLimitMiddleware(ipLimiter)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	testServer := httptest.NewServer(middleware(handler))
	defer testServer.Close()

	client := testServer.Client()

	// First request should pass
	resp, err := client.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, resp.StatusCode)
	}

	// Second request should be rate limited
	resp, err = client.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status %v, got %v", http.StatusTooManyRequests, resp.StatusCode)
	}
}
