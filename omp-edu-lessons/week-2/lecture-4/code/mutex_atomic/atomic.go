package mutex_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicIntro() {
	// Введение в atomic
	m := sync.Mutex{}
	// Смотрим что под капотом
	m.Lock()
	m.Unlock()
	// -----
	// Пробуем сохранить/вытащить значение
	at := atomic.Value{}
	at.Store(100)
	ct := at.Load()
	fmt.Println(ct)
}
