package main

import (
	"fmt"
	"time"
)
// Имитатор WaitGroup
// Здесь куча проблем внутри
type wait struct {
	count uint
	ch chan interface{}
}

func initWait(count uint) *wait {
	return &wait{count: count, ch: make(chan interface{}, count)}
}

func (w *wait)Done(){
	defer func(){
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
		w.ch <- 1
}

func (w *wait)Wait() {
	defer close(w.ch)
	for {
		<- w.ch
		w.count--
		fmt.Println("minus")
		if w.count <= 0 {
			return
		}
	}
}

func test(id int, w *wait) {
	defer w.Done()
	fmt.Printf("<%v> done\n", id)
}

func main(){
	w := initWait(5)

	for i := 0; i < 5; i++ {
		go test(i, w)
	}

	go func(){
		w.Wait()
		fmt.Println("second done")
	}()
	time.Sleep(time.Second)
	w.Wait()
	time.Sleep(time.Second)
	fmt.Println("Done")
}
