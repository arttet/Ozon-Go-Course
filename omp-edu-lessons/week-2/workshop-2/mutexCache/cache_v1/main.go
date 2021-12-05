package cache_v1

import (
	"time"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/helpers"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/storage"
)

type simpleCache struct {
	storage map[string]string
}

func NewCache() storage.Cache {
	return &simpleCache{
		storage: make(map[string]string),
	}
}

func (c *simpleCache) Set(key, value string) error {
	c.storage[key] = value

	return nil
}

func (c *simpleCache) get(key string) (string, error) {
	time.Sleep(time.Second)
	value, ok := c.storage[key]
	if !ok {
		return "", helpers.ErrNotFound
	}

	return value, nil
}

func (c *simpleCache) Get(key string) (string, error) {
	return c.get(key)
}

func (c *simpleCache) Delete(key string) error {
	delete(c.storage, key)

	return nil
}
