package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Паттерн / будет обрабатывать все запросы
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "Default handler")
	})

	mux.HandleFunc("/1", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "One")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
