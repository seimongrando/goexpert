package strategy

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go/wait"
)

func TestRedisStore(t *testing.T) {
	// Configuração do Test Container
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:7.0-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp"),
	}
	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("failed to start Redis container: %v", err)
	}
	defer redisContainer.Terminate(ctx) // Garante que o container será desligado após o teste

	// Obtém o endereço do Redis
	host, err := redisContainer.Host(ctx)
	if err != nil {
		t.Fatalf("failed to get Redis container host: %v", err)
	}
	port, err := redisContainer.MappedPort(ctx, "6379/tcp")
	if err != nil {
		t.Fatalf("failed to get Redis container port: %v", err)
	}
	addr := host + ":" + port.Port()

	// Configuração do RedisStore
	store := NewRedisStore(addr, "", 0)

	key := "test-key"
	expiration := 2 * time.Second

	// Incrementa a chave
	count, err := store.Increment(key, expiration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count to be 1, got %d", count)
	}

	// Incrementa novamente
	count, err = store.Increment(key, expiration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count to be 2, got %d", count)
	}

	// Aguarda expiração
	time.Sleep(3 * time.Second)
	count, err = store.Increment(key, expiration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count to reset to 1 after expiration, got %d", count)
	}

	// Reseta a chave
	store.Reset(key)
	count, err = store.Increment(key, expiration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count to reset to 1 after reset, got %d", count)
	}
}
