package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	router          *mux.Router
	productProvider ProductProvider
}

// ProductProvider отдает товар по SKU
type ProductProvider interface {
	GetProduct(ctx context.Context, sku int64) (*Product, error)
}

func newHandler(productProvider ProductProvider) *httpHandler {
	router := mux.NewRouter()
	handler := &httpHandler{
		router:          router,
		productProvider: productProvider,
	}
	router.HandleFunc("/products/{sku:[0-9]+}", handler.getProduct).Methods(http.MethodGet)
	return handler
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *httpHandler) getProduct(w http.ResponseWriter, req *http.Request) {
	// Получим SKU товара из запроса
	vars := mux.Vars(req)
	sku, _ := strconv.ParseInt(vars["sku"], 10, 64)
	if sku <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.productProvider.GetProduct(req.Context(), sku)
	if errors.Is(err, ErrProductNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(result)
}
