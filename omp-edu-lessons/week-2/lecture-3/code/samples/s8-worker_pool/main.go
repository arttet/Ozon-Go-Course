package main

import (
	"fmt"
	"time"
)

type Pool struct {
	in chan int
	out chan string
	runners uint
	task func(int) string
}

func NewPool(buffer uint) *Pool {
	return &Pool{
		in: make(chan int),
		out: make(chan string, buffer),
		runners: buffer,
	}
}

func (p *Pool)SetTask(task func(int)string) {
	p.task = task
}

func (p *Pool)Run(){
	counter := make(chan interface{}, p.runners)
	for {
			in, ok := <-p.in
			if !ok {
				return
			}
			x := true
			for x {
				select {
					case counter <- 0:
					x = false
				default:
					fmt.Printf("%v\twaiting\n", in)
					time.Sleep(time.Millisecond * 100)
				}
			}
			go func(in int) {
				defer func() {
					<-counter
				}()
				p.out <- p.task(in)
			}(in)
	}
}

func (p *Pool)OutChan() <-chan string {
	return p.out
}

func (p *Pool)Send(s int) {
	p.in <- s
}

func main(){
	p := NewPool(10)
	p.SetTask(func(s int)string{
		t := "low"
		if s % 2 == 0 {
			time.Sleep(time.Second)
		} else {
			t = "fast"
			time.Sleep(time.Millisecond * 300)
		}
		time.Sleep(time.Second)
		return fmt.Sprintf("%v\t%v handled", s, t)
	})

	go func(){
		for x := range p.OutChan() {
			fmt.Println(x)
		}
	}()

	go func() {
		for i := 0; i< 1000; i++ {
			go p.Send(i)
			time.Sleep(time.Millisecond*100)
		}
	}()

	go p.Run()

	time.Sleep(time.Second * 10)
}