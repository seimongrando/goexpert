package limiter

import "time"

type rateLimiter struct {
	store      store
	ipLimit    int
	tokenLimit int
	blockTime  time.Duration
}
