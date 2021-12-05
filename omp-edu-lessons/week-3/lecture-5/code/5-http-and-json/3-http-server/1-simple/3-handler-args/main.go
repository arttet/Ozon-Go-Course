package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "%+v\n", req)

		// Можно без форматирования строки
		_, _ = fmt.Fprintln(w, "Text 1")

		// Можно напрямую вызвать метод в интерфейсе http.ResponseWriter
		_, _ = w.Write([]byte("Text 2\n"))
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
