package main

import "fmt"

type Result struct {
	ID int
	Result int
}

func consumer(id int, in <-chan int, out chan<- Result) {
	res := 0
	for x := range in {
		res += x
	}
	out<- Result{ID: id, Result: res}
}

func main() {
	inch := make(chan int)
	outch := make(chan Result)

	for i := 0; i < 10; i++ {
		go consumer(i, inch, outch)
	}

	for i:=0; i < 1000; i++ {
		inch <- i
	}

	close(inch)

	for i := 0; i < 10; i++ {
		res, ok := <-outch
		if !ok {
			panic("Out channel is closed")
		}
		fmt.Printf("# %v: %v\n", res.ID, res.Result)
	}

	close(outch)
}
