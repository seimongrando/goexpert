package main

import (
	"net/http"
	"net/http/httptest"
	"rate-limiter/limiter"
	"rate-limiter/middleware"
	"rate-limiter/strategy"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	// Configuração do RateLimiter com MemoryStore
	store := strategy.NewMemoryStore()
	rl := limiter.NewRateLimiter(store, 2, 5, 1*time.Second)

	// Configuração do middleware
	mux := http.NewServeMux()
	mux.Handle("/", middleware.RateLimiterMiddleware(rl)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request successful"))
	})))

	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	// Cenário 1: Requisição válida (IP não limitado)
	req, _ := http.NewRequest("GET", server.URL, nil)
	req.RemoteAddr = "192.168.0.1"
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	// Cenário 2: Excedendo o limite de IP
	for i := 0; i < 3; i++ { // Limite é 2, o terceiro deve falhar
		resp, err = client.Do(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", resp.StatusCode)
	}
}
