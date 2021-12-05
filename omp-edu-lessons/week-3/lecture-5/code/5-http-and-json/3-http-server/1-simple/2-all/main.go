package main

import (
	"fmt"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello, World!")
	}

	http.HandleFunc("/", helloHandler)

	_ = http.ListenAndServe(":8080", nil)
	// Или вот так:
	// _ = http.ListenAndServe("0.0.0.0:8080", nil)
}
