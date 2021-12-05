package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func createSwaggerServer(gatewayAddr, swaggerAddr, swaggerPath string) (*http.Server, error) {
	originalSwagger, err := os.ReadFile(swaggerPath)
	if err != nil {
		return nil, fmt.Errorf("missing swagger.json: %w", err)
	}

	patchedSwagger, err := injectHost(originalSwagger, gatewayAddr)
	if err != nil {
		return nil, err
	}

	serveMux := http.NewServeMux()

	serveMux.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write(patchedSwagger)
	}))

	docsServer := http.FileServer(http.Dir("./swagger/dist"))

	serveMux.Handle("/docs/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/docs" || strings.HasPrefix(r.URL.Path, "/docs/") {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/docs")
			docsServer.ServeHTTP(w, r)
		} else {
			w.WriteHeader(404)
		}
	}))

	gatewayServer := &http.Server{
		Addr:    swaggerAddr,
		Handler: serveMux,
	}

	return gatewayServer, nil
}

func injectHost(swaggerBytes []byte, host string) ([]byte, error) {
	parsedSwagger := map[string]interface{}{}
	err := json.Unmarshal(swaggerBytes, &parsedSwagger)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}
	parsedSwagger["host"] = host

	resultBytes, err := json.Marshal(parsedSwagger)
	if err != nil {
		return nil, fmt.Errorf("unexpected json marshal error: %w", err)
	}

	return resultBytes, nil
}
