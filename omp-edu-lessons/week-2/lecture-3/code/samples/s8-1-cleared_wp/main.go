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
		// получаем данные для обработки
		in, ok := <-p.in
		if !ok {
			return
		}
		// Проеверяем, можно ли запустить еще один воркер. Если в очереди есть место, то воркеров меньше максимума
		// Записываем в очередь что-нибудь, чтобы занять слот
		counter <- 0
		// Выполняем задачу в отдельной горутине
		go func(in int) {
			defer func() {
				// Читаем из очереди, чтобы освободить слот
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

func taskForRunners(s int)string{
	t := "low"
	if s % 2 == 0 {
		time.Sleep(time.Second)
	} else {
		t = "fast"
		time.Sleep(time.Millisecond * 300)
	}
	time.Sleep(time.Second)
	return fmt.Sprintf("%v\t%v handled", s, t)
}

func main(){
	p := NewPool(10)
	p.SetTask(taskForRunners)

	go func(){
		// Читаем результаты работы из канала вывода и печатаем результаты
		for x := range p.OutChan() {
			fmt.Println(x)
		}
	}()

	go func() {
		for i := 0; i< 1000; i++ {
			// Отправляем данные для обработки во входной канал
			go p.Send(i)
			time.Sleep(time.Millisecond*100)
		}
	}()

	// Запускаем воркеры
	go p.Run()

	time.Sleep(time.Second * 10)
}