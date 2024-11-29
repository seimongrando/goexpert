package strategy

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// NewRedisStore cria uma nova instância do redisStore
func NewRedisStore(addr, password string, db int) *redisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisStore{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Increment incrementa o contador associado à chave no Redis
func (r *redisStore) Increment(key string, expiration time.Duration) (int, error) {
	// Incrementa o valor no Redis
	count, err := r.client.Incr(r.ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to increment key %s: %w", key, err)
	}

	// Define a expiração se for uma nova chave ou reinicia a expiração
	if count == 1 {
		err = r.client.Expire(r.ctx, key, expiration).Err()
		if err != nil {
			return 0, fmt.Errorf("failed to set expiration for key %s: %w", key, err)
		}
	}

	return int(count), nil
}

// Reset remove a chave do Redis
func (r *redisStore) Reset(key string) {
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		fmt.Printf("failed to reset key %s: %v\n", key, err)
	}
}
