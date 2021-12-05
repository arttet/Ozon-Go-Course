package context

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func contextDone() {
	// Context с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	err := gracefulWorker(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

// gracefulWorker обращает внимание на отмены контекста ctx.Done()
func gracefulWorker(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Second):
		return errors.New("worker error")
	}
}
