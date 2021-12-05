package main

import (
	"log"
	"net/http"
)

type httpHandler struct {
	mux *http.ServeMux
}

func newHandler() *httpHandler {
	handler := &httpHandler{
		mux: http.NewServeMux(),
	}
	handler.mux.HandleFunc("/", handler.index)
	return handler
}

// ServeHTTP просто перенаправляет запрос в *http.ServeMux
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (httpHandler) index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("index"))
}

func main() {
	mux := newHandler()
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
