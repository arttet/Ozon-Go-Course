package main

import (
	"fmt"
	"time"
)

func main(){
	//runtime.GOMAXPROCS(1)

	ch := make([]chan int, 3)
	ch[0] = make(chan int)
	ch[1] = make(chan int)
	ch[2] = make(chan int)

	for i:= 0; i < 3; i++ {
		go func(i int) {
			for y := 0; y < 5; y++ {
				ch[2-i] <- y
			}
		}(i)
	}

	go func() {
		for {
			select {
			case x := <-ch[0]:
				fmt.Printf("<%v> -> %v\n", 0, x)
			case x := <-ch[1]:
				fmt.Printf("<%v> -> %v\n", 1, x)
			case x := <-ch[2]:
				fmt.Printf("<%v> -> %v\n", 2, x)
			}
		}
	}()

	time.Sleep(time.Second)
}
