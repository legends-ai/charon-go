package riot

import (
	"sync"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/time/rate"
)

// RateLimiter rate limits.
type RateLimiter struct {
	isLimited   bool
	isLimitedMu sync.RWMutex
	rl          *rate.Limiter
}

// NewRateLimiter creates a new RateLimiter.
func NewRateLimiter(rateLimit time.Duration) *RateLimiter {
	r := &RateLimiter{
		rl: rate.NewLimiter(rate.Every(rateLimit), 1),
	}
	return r
}

// Wait waits until the rate limiter is ready.
func (r *RateLimiter) Wait(ctx context.Context) {
	for {
		r.rl.Wait(ctx)

		// check if we have a retry after going on
		r.isLimitedMu.RLock()
		limited := r.isLimited
		r.isLimitedMu.RUnlock()
		if !limited {
			break
		}
	}
}

// RetryAfter makes all requests wait until RetryAfter is over.
func (r *RateLimiter) RetryAfter(d time.Duration) {
	r.isLimitedMu.Lock()
	r.isLimited = true
	r.isLimitedMu.Unlock()

	go func() {
		time.Sleep(d)
		r.isLimitedMu.Lock()
		r.isLimited = false
		r.isLimitedMu.Unlock()
	}()
}
