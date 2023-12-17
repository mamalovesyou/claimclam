package ratelimiter_test

import (
	"testing"

	"github.com/mamalovesyou/claimclam/internal/ratelimiter"
)

func TestAddAndAllow(t *testing.T) {
	ipLimiter := ratelimiter.NewIPRateLimiter(1, 5)
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
