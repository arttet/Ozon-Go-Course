package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/gorilla/mux"
)

type httpPrivateHandler struct {
	router *mux.Router
}

func newPrivateHandler() *httpPrivateHandler {
	router := mux.NewRouter()
	handler := &httpPrivateHandler{router: router}
	router.HandleFunc("/version", handler.version).Methods(http.MethodGet)
	router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	return handler
}

func (h *httpPrivateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (httpPrivateHandler) version(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	const versionFormat = `{"go_version": %q, "goarch": %q, "goos": %q}`
	_, _ = fmt.Fprintf(w, versionFormat, runtime.Version(), runtime.GOARCH, runtime.GOOS)
}
