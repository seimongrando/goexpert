package limiter

import "time"

func NewRateLimiter(store store, ipLimit, tokenLimit int, blockTime time.Duration) *rateLimiter {
	return &rateLimiter{
		store:      store,
		ipLimit:    ipLimit,
		tokenLimit: tokenLimit,
		blockTime:  blockTime,
	}
}

func (rl *rateLimiter) IsRateLimited(key string, isToken bool) bool {
	// Determina o limite baseado no tipo de chave
	limit := rl.ipLimit
	if isToken {
		limit = rl.tokenLimit
	}

	// Incrementa e verifica o limite
	count, err := rl.store.Increment(key, rl.blockTime)
	if err != nil {
		return false
	}
	return count > limit
}
