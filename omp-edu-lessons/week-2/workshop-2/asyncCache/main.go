package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	fmt.Printf("1 len %v\n", len(ch))
	go func() {
		ch <- "1111"
		fmt.Println("added")
	}()
	time.Sleep(time.Second)
	fmt.Printf("2 len %v\n", len(ch))
	a, ok := <- ch
	fmt.Printf("%v %v\n", a, ok)
}
