package limiter

import (
	"rate-limiter/strategy"
	"testing"
	"time"
)

func TestRateLimiter_IsRateLimited(t *testing.T) {
	store := strategy.NewMemoryStore()
	rl := NewRateLimiter(store, 2, 5, 1*time.Second)

	// Testa limite de IP
	ipKey := "192.168.0.1"
	if rl.IsRateLimited(ipKey, false) {
		t.Error("expected IP to not be rate limited initially")
	}

	rl.IsRateLimited(ipKey, false) // 1ª requisição
	rl.IsRateLimited(ipKey, false) // 2ª requisição
	if !rl.IsRateLimited(ipKey, false) {
		t.Error("expected IP to be rate limited after exceeding limit")
	}

	// Testa limite de Token
	tokenKey := "test-token"
	if rl.IsRateLimited(tokenKey, true) {
		t.Error("expected token to not be rate limited initially")
	}

	for i := 0; i < 5; i++ {
		rl.IsRateLimited(tokenKey, true) // 5 requisições
	}
	if !rl.IsRateLimited(tokenKey, true) {
		t.Error("expected token to be rate limited after exceeding limit")
	}
}
