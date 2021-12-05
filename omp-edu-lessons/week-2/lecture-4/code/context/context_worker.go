package context

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func contextWorker() {
	// Context с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)

	go func() {
		err := worker2(ctx)
		if err != nil {
			fmt.Println(err)
			cancel()
		}
	}()

	time.Sleep(time.Second * 2)
}

// Эмулирует ошибку + полезную работу
func worker2(ctx context.Context) error {
	time.Sleep(time.Second)
	return errors.New("worker error")
}
