package strategy

import (
	"testing"
	"time"
)

func TestMemoryStore_Increment(t *testing.T) {
	store := NewMemoryStore()

	// Incrementa uma chave e verifica o valor
	key := "test-key"
	count, err := store.Increment(key, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count to be 1, got %d", count)
	}

	// Incrementa novamente e verifica o valor atualizado
	count, err = store.Increment(key, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count to be 2, got %d", count)
	}

	// Aguarda expiração
	time.Sleep(3 * time.Second)
	count, err = store.Increment(key, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count to reset to 1 after expiration, got %d", count)
	}
}

func TestMemoryStore_Reset(t *testing.T) {
	store := NewMemoryStore()

	// Incrementa uma chave
	key := "test-key"
	_, _ = store.Increment(key, 2*time.Second)

	// Reseta a chave
	store.Reset(key)

	// Incrementa novamente e verifica se começa de 1
	count, err := store.Increment(key, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count to reset to 1, got %d", count)
	}
}
