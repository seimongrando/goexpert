package limiter

import "time"

// store define a interface para o armazenamento do Rate Limiter
type store interface {
	Increment(key string, expiration time.Duration) (int, error)
	Reset(key string)
}
