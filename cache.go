package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrorNotFound error = errors.New("value not found")
	ErrorExpired  error = errors.New("value is expired")
)

type item struct {
	value       interface{}
	dedlineTime time.Time
}

type cache struct {
	storage map[string]item
	mu      sync.RWMutex
}

func New() *cache {
	return &cache{
		storage: make(map[string]item),
	}
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.storage[key] = item{
		value:       value,
		dedlineTime: time.Now().Add(ttl),
	}
}

func (c *cache) Get(key string) (interface{}, error) {
	c.mu.RLock()

	value, ok := c.storage[key]

	c.mu.RUnlock()

	if !ok {
		return nil, ErrorNotFound
	}

	if value.dedlineTime.Before(time.Now()) {
		return nil, ErrorExpired
	}

	return value, nil
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}
