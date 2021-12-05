package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/heart", func(w http.ResponseWriter, req *http.Request) {
		// Без этой строчки будет неправильный Content-Type
		w.Header().Set("Content-Type", "application/json")

		_, _ = fmt.Fprintln(w, `{"heart_rate": 80}`)
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
