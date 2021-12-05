package no_mutex

import (
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/helpers"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/storage"
)

// Имплементация кеша без мютексов
type impl struct {
	st map[string]string
}

func New() storage.Cache {
	return &impl{st: map[string]string{}}
}

func (i *impl) Get(key string) (string, error) {
	v, ok := i.st[key]
	if !ok {
		return "", helpers.ErrNotFound
	}

	return v, nil
}

func (i *impl) Set(key, value string) error {
	i.st[key] = value
	return nil
}

func (i *impl) Delete(key string) error {
	delete(i.st, key)
	return nil
}
