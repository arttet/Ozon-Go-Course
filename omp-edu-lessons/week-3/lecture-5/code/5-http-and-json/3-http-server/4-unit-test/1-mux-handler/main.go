package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello World"))
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	return router
}
