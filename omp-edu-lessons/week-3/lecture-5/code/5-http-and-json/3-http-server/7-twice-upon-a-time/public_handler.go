package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type httpPublicHandler struct {
	router *mux.Router
}

func newPublicHandler() *httpPublicHandler {
	handler := &httpPublicHandler{
		router: mux.NewRouter(),
	}
	handler.router.HandleFunc("/index", handler.index).Methods(http.MethodGet)
	return handler
}

func (h *httpPublicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (httpPublicHandler) index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Public Index")
}
