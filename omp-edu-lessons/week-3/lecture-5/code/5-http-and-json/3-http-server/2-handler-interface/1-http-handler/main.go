package main

import (
	"fmt"
	"log"
	"net/http"
)

type HttpHandler struct{}

// ServeHTTP расширяет структуру HttpHandler реализуя интерфейс http.Handler
func (HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Был запрошен путь: %q", req.RequestURI)
}

func main() {
	const addr = "127.0.0.1:8080"
	handler := HttpHandler{}

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
	// Тоже самое что и
	// log.Fatal(http.ListenAndServe(addr, handler))
}
