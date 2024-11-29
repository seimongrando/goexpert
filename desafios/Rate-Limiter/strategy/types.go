package strategy

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type memoryStore struct {
	data map[string]*entry
	mu   sync.Mutex
}

type entry struct {
	count      int
	expiration time.Time
}

type redisStore struct {
	client *redis.Client
	ctx    context.Context
}
