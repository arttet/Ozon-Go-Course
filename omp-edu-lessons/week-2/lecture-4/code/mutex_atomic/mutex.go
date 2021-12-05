package mutex_atomic

import (
	"fmt"
	"sync"
	"time"
)

// 1 запуск, concurrent map writes
func concurrentWrites() {
	cs := map[string]int{"касса": 0}

	for i := 0; i < 1000; i++ {
		go func(k int) {
			cs["касса"] += 1
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}

// 2 запуск, работает как задуманно через mutex
func mutexSync() {
	cs := map[string]int{"касса": 0}
	m := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func(k int) {
			m.Lock()
			defer m.Unlock()
			cs["касса"] += 1
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
