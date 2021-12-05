package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	publicServer := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: newPublicHandler(),
	}
	go func() {
		srvErr := publicServer.ListenAndServe()
		if srvErr != http.ErrServerClosed {
			log.Printf("publicServer.ListenAndServe(): %v\n", srvErr)
		}

		fmt.Println("publicServer successfully stopped")
	}()

	privateServer := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: newPrivateHandler(),
	}
	go func() {
		srvErr := privateServer.ListenAndServe()
		if srvErr != http.ErrServerClosed {
			log.Fatalf("privateServer.ListenAndServe(): %v", srvErr)
		}

		fmt.Println("privateServer successfully stopped")
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := publicServer.Shutdown(closeCtx)
		if err != nil {
			log.Println("publicServer.Shutdown() failed:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := privateServer.Shutdown(closeCtx)
		if err != nil {
			log.Println("privateServer.Shutdown() failed:", err)
		}
	}()

	wg.Wait()
}
