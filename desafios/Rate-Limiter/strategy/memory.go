package strategy

import (
	"time"
)

// NewMemoryStore cria uma nova instância do armazenamento em memória
func NewMemoryStore() *memoryStore {
	return &memoryStore{
		data: make(map[string]*entry),
	}
}

func (m *memoryStore) Increment(key string, expiration time.Duration) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	if e, exists := m.data[key]; exists {
		// Verifica se o tempo expirou
		if now.After(e.expiration) {
			// Reinicia a contagem se expirou
			e.count = 1
			e.expiration = now.Add(expiration)
		} else {
			e.count++
		}
		return e.count, nil
	}

	// Nova entrada
	m.data[key] = &entry{
		count:      1,
		expiration: now.Add(expiration),
	}
	return 1, nil
}

func (m *memoryStore) Reset(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}
