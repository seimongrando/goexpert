package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockLimiter é uma implementação mockada do RateLimiter para fins de teste.
type MockLimiter struct {
	rateLimited map[string]bool
}

func NewMockLimiter() *MockLimiter {
	return &MockLimiter{
		rateLimited: make(map[string]bool),
	}
}

func (m *MockLimiter) IsRateLimited(key string, isToken bool) bool {
	return m.rateLimited[key]
}

func (m *MockLimiter) SetRateLimited(key string, limited bool) {
	m.rateLimited[key] = limited
}

func TestRateLimiterMiddleware(t *testing.T) {
	// Mock do limiter
	mockLimiter := NewMockLimiter()

	// Middleware para teste
	middleware := RateLimiterMiddleware(mockLimiter)

	// Handler de teste
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	// Cenário 1: Requisição permitida (IP não limitado)
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.0.1:12345" // Inclui porta no RemoteAddr
	rec := httptest.NewRecorder()

	mockLimiter.SetRateLimited("192.168.0.1", false)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	// Cenário 2: Requisição bloqueada (IP limitado)
	rec = httptest.NewRecorder()
	mockLimiter.SetRateLimited("192.168.0.1", true)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", rec.Code)
	}

	// Cenário 3: Requisição permitida com token
	req.Header.Set("API_KEY", "token123")
	rec = httptest.NewRecorder()
	mockLimiter.SetRateLimited("token123", false)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	// Cenário 4: Requisição bloqueada com token
	rec = httptest.NewRecorder()
	mockLimiter.SetRateLimited("token123", true)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", rec.Code)
	}

	// Cenário 5: Endereço IP inválido
	req.RemoteAddr = "invalid-addr"
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for invalid IP address, got %d", rec.Code)
	}
}
