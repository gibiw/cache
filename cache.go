package cache

import (
	"sync"
)

type cache struct {
	storage map[string]interface{}
	mu      sync.Mutex
}

func New() *cache {
	return &cache{
		storage: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.storage[key]; ok {
		c.storage[key] = value
		return
	}

	c.storage[key] = value
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, ok := c.storage[key]; ok {
		return value, true
	}

	return nil, false
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}
