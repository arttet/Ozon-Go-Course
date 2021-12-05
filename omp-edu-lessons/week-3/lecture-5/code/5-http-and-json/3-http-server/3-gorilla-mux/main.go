package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	router *mux.Router
}

func newHandler() *httpHandler {
	router := mux.NewRouter()
	handler := &httpHandler{router: router}
	router.HandleFunc("/index", handler.index).Methods(http.MethodGet)
	return handler
}

// ServeHTTP просто перенаправляет запрос в *mux.Router
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (httpHandler) index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "ok")
}

func main() {
	handler := newHandler()
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", handler))
}
