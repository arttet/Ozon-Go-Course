package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	notify := make(chan os.Signal, 1)
	signal.Notify(notify, syscall.SIGHUP)

	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT)

	for {
		select {
		case <-notify:
			fmt.Println("Notify signal")
		case x := <- kill:
			fmt.Printf("Kill signal %#v\n", x)
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println("ups...")
		}
	}
}
