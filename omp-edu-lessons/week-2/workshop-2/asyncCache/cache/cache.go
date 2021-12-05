package cache

import (
	"fmt"
	"sync"
	"time"
)

const timeout = time.Millisecond * 2

type Cache struct {
	c map[string]string
	m *sync.RWMutex
}

func InitCache() *Cache {
	return &Cache{
		c: make(map[string]string),
		m: new(sync.RWMutex),
	}
}

func (c *Cache) Add(k, v string) {
	c.m.Lock()
	defer c.m.Unlock()

	time.Sleep(timeout)
	c.c[k] = v
}

func (c *Cache) Get(k string) (string, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	time.Sleep(timeout)
	v, ok := c.c[k]
	return v, ok
}

func (c *Cache) Del(k string) {
	c.m.Lock()
	defer c.m.Unlock()

	time.Sleep(timeout)
	delete(c.c, k)
}

func (c *Cache)String()string {
	return fmt.Sprintf("%#v", c.c)
}