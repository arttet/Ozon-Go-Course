package bench

import (
	"sync"
	"sync/atomic"
	"time"
)

func MutexCounter() int {
	// Количество выполняемых Goroutine
	goroutinesCount := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	// Запускаем горетины и увеличиваем счетчик
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			goroutinesCount++
			m.Unlock()
			// Имитируем полезную работу
			time.Sleep(time.Microsecond)
			m.Lock()
			goroutinesCount--
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCount
}

func AtomicCounter() int32 {
	// Количество выполняемых Goroutines
	goroutinesCount := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCount, 1)
			// Имтируем полезную работу
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&goroutinesCount, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCount
}
