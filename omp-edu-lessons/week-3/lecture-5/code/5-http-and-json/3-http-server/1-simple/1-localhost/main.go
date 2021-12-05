package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello, World!")
	}

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
