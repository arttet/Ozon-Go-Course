package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for x:= 0; x < 10; x++ {
		go func() {
			fmt.Printf("goroutine %v\n", x)
		}()
		runtime.Gosched()
	}

	time.Sleep(time.Second * 1)
}
