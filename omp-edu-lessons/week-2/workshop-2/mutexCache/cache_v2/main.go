package cache_v2

import (
	"sync"
)

type SimpleCache struct {
	m  map[string]string
	mu sync.RWMutex
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{m: make(map[string]string)}
}

func (c *SimpleCache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
	return nil
}

func (c *SimpleCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
	return nil
}
