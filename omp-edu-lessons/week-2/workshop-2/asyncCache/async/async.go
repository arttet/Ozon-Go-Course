package async

import (
	"context"
	"errors"

	"github.com/kshmatov/asyncCache/cache"
)

var ErrTimeout = errors.New("timeout")

type Cache struct {
	c *cache.Cache
}

func InitAsyncCache() *Cache {
	return &Cache{
		c: cache.InitCache(),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		v, ok := c.c.Get(key)
		if ok {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return "", ErrTimeout
	case x, ok := <-ch:
		if ok {
			return x, nil
		}
		return "", errors.New("not found")
	}
}

func (c *Cache) Add(ctx context.Context, key, value string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Add(key, value)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case <-ch:
		return nil
	}
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Del(key)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case <-ch:
		return nil
	}
}

