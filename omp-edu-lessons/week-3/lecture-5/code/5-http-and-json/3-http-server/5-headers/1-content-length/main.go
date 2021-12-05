package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/heart", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "❤️")
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
