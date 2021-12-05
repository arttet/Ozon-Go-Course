package with_mutex

import (
	"sync"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/helpers"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/storage"
)

// Имплементация кеша с mutex
type impl struct {
	st map[string]string
	mu sync.Mutex
}

func New() storage.Cache {
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
	return nil
}

func (i *impl) Delete(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.st, key)
	return nil
}
