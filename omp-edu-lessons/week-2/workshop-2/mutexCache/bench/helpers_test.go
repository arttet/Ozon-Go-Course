package bench

import (
	"errors"
	"fmt"
	"sync"

	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/helpers"
	"gitlab.ozon.ru/vserdyukov/go-kafedra/week2-workshop/storage"
)

// emulateLoad вспомогательная функция, создает нагрузку на кеш через горутины
func emulateLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		// С этими key/value будем работать на этой итерации цикла
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)
		// Запись в кеш
		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

		wg.Add(1)
		// Чтение из кеша
		go func(k, v string) {
			_, err := c.Get(k)
			// Проверим, что ошибка не связана с тем что записи нет в кеше
			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		// Удаление из кеша
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)
	}

	// Ждем пока все горутины отработают
	wg.Wait()
}

// emulateLoadWithMetrics вспомогательная функция, создает нагрузку на кеш через горутины и проверяет количество записей в кеше
func emulateLoadWithMetrics(cm storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		// CacheWithMetrics также реализует интерфейс Cache
		// По этому работает как есть
		emulateLoad(cm, parallelFactor)
		wg.Done()
	}()

	// Добавим забор метрик с кеша
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			_ = cm.TotalAmount()
			wg.Done()
		}()
	}

	wg.Wait()
}

// emulateLoad вспомогательная функция, создает нагрузку на кеш через горутины
func emulateReadIntensiveLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	// Понижаем в 10 раз нагрузку на запись и удаление
	for i := 0; i < parallelFactor/10; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)
		// Запись в кеш
		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

		wg.Add(1)
		// Удаление из кеша
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)
	}

	// Нагрузку на чтение оставляем как есть
	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)
		// Чтение из кеша
		go func(k, v string) {
			_, err := c.Get(k)
			// Проверим, что ошибка не связана с тем что записи нет в кеше
			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)

	}

	// Ждем пока все горутины отработают
	wg.Wait()
}
