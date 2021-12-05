package with_atomic

import (
	"sync"
	"sync/atomic"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/helpers"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/storage"
)

// // Имплементация кеша с mutex и atomic метриками
type impl struct {
	st    map[string]string
	mu    sync.Mutex
	total int64
}

func New() storage.CacheWithMetrics {
	return &impl{
		st: map[string]string{},
		mu: sync.Mutex{},
	}
}

func (i *impl) Get(key string) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, ok := i.st[key]
	if !ok {
		return "", helpers.ErrNotFound
	}

	return v, nil
}

func (i *impl) Set(key, value string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.st[key] = value
	atomic.StoreInt64(&i.total, 1)
	return nil
}

func (i *impl) Delete(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.st, key)
	atomic.StoreInt64(&i.total, -1)
	return nil
}

func (i *impl) TotalAmount() int64 {
	return atomic.LoadInt64(&i.total)
}
