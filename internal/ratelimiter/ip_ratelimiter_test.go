package ratelimiter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAndAllow(t *testing.T) {
	ipLimiter := NewIPRateLimiter(1, 5)
	ip := "127.0.0.1"

	// Test AddIP
	limiter := ipLimiter.AddIP(ip)
	if limiter == nil {
		t.Errorf("AddIP() failed, expected limiter, got nil")
	}

	// Test GetLimiter for existing IP
	ipAllowed := ipLimiter.Allow(ip)
	if ipAllowed != true {
		t.Errorf("Allow() failed, expected true, got false")
	}

	// Test GetLimiter for new IP
	newIP := "127.0.0.2"
	newIPAllowed := ipLimiter.Allow(newIP)
	if newIPAllowed != true {
		t.Errorf("Allow() for new IP failed, expected true, got false")
	}
}

func TestIPRateLimitMiddleware(t *testing.T) {
	ipLimiter := NewIPRateLimiter(1, 1) // Set low limit for testing
	middleware := IPRateLimitMiddleware(ipLimiter)

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
