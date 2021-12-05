package channel

import (
	"errors"

	"github.com/kshmatov/asyncCache/cache"
)

var ErrNotFound = errors.New("not found")

type Data struct {
	Key string
	Value string
}

type ChanCache struct {
	c *cache.Cache
	inChan chan Data
}

func Init() *ChanCache {
	return &ChanCache{
		c: cache.InitCache(),
		inChan: make(chan Data),
	}
}

func (c *ChanCache)GetChannel()chan <- Data {
	return c.inChan
}

func (c *ChanCache)Run(){
	go func(){
		for {
			x, ok := <- c.inChan
			if !ok {
				return
			}
			c.c.Add(x.Key, x.Value)
		}
	}()
}

func (c *ChanCache)Get(key string) (string, error) {
	v, ok := c.c.Get(key)
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (c *ChanCache)ChannelGet(key string) <-chan string {
	ch := make(chan string)
	go func(){
		defer close(ch)
		v, ok := c.c.Get(key)
		if ok {
			ch <- v
		}
	}()
	return ch
}
