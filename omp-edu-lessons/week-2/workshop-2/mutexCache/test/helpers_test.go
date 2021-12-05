package test

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/helpers"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/storage"
)

// emulateLoad вспомогательная функция, создает нагрузку на кеш через горутины
func emulateLoad(t *testing.T, c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		// С этими key/value будем работать на этой итерации цикла
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)
		// Запись в кеш
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key)

		wg.Add(1)
		// Чтение из кеша
		go func(k, v string) {
			storedValue, err := c.Get(k)
			// Если другая горутина не успела удалить значение из кеша
			// Проверим, что оно совпадает с тем что мы хотели добавить в кеш
			if !errors.Is(err, helpers.ErrNotFound) {
				assert.Equal(t, v, storedValue)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		// Удаление из кеша
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)
	}

	// Ждем пока все горутины отработают
	wg.Wait()
}

// emulateLoadWithMetrics вспомогательная функция, создает нагрузку на кеш через горутины и проверяет количество записей в кеше
func emulateLoadWithMetrics(t *testing.T, cm storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		// CacheWithMetrics также реализует интерфейс Cache
		// По этому работает как есть
		emulateLoad(t, cm, parallelFactor)
		wg.Done()
	}()

	// Добавим забор метрик с кеша
	var min, max int64
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			total := cm.TotalAmount()
			if total > max {
				max = total
			}
			if total < min {
				min = total
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(max, min)
}
