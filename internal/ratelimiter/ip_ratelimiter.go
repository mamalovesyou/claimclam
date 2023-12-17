package ratelimiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter manages rate limiters for individual IP addresses.
type IPRateLimiter struct {
	ips map[string]*rate.Limiter // Map to store rate limiters per IP address
	mu  *sync.RWMutex            // Mutex to synchronize access to the map
	r   rate.Limit               // Rate of requests allowed
	b   int                      // Burst size allowed
}

// NewIPRateLimiter creates a new instance of IPRateLimiter.
// It takes a rate (r) which specifies the number of requests allowed per second,
// and a burst size (b) which specifies the maximum burst size allowed.
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter for a given IP address and adds it to the ips map
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

// Allow returns true if the provided IP address is allowed by the rate limiter.
func (i *IPRateLimiter) Allow(ip string) bool {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	i.mu.Unlock()
	if !exists {
		limiter = i.AddIP(ip)
	}
	return limiter.Allow()
}
