package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Паттерн / будет обрабатывать все запросы
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "Default handler")
	})

	http.HandleFunc("/1", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "One")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
