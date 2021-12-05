package helpers

import (
	"errors"
)

var ErrNotFound = errors.New("value not found")

// 1 реализация
/*
type simpleCache struct {
	storage map[string]string
}

var ErrNotFound = errors.New("value not found")

func NewCache() Cache {
	return &simpleCache{
		storage: make(map[string]string),
	}
}

func (c *simpleCache) Set(key, value string) error {
	c.storage[key] = value

	return nil
}

func (c *simpleCache) Get(key string) (string, error) {
	value, ok := c.storage[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (c *simpleCache) Delete(key string) error {
	delete(c.storage, key)

	return nil
}
*/

// 2 реализация
//type SimpleCache struct {
//	m  map[string]string
//	mu sync.RWMutex
//}
//
//func NewSimpleCache() *SimpleCache {
//	return &SimpleCache{m: make(map[string]string)}
//}
//
//func (c *SimpleCache) Get(key string) (string, error) {
//	c.mu.RLock()
//	defer c.mu.RUnlock()
//	value, ok := c.m[key]
//	if !ok {
//		return "", ErrNotFound
//	}
//	return value, nil
//}
//
//func (c *SimpleCache) Set(key, value string) error {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	c.m[key] = value
//	return nil
//}
//
//func (c *SimpleCache) Delete(key string) error {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	delete(c.m, key)
//	return nil
//}
