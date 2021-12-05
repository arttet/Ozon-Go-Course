package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type productProviderMock struct {
	validSKU int64
}

func (m *productProviderMock) GetProduct(_ context.Context, sku int64) (*Product, error) {
	if sku == m.validSKU {
		return &Product{SKU: sku, Name: "Valid product"}, nil
	}
	return nil, ErrProductNotFound
}

func Test_httpHandler_getProduct(t *testing.T) {
	t.Run("invalid sku", func(t *testing.T) {
		// Arrange
		productProvider := &productProviderMock{}
		mux := newHandler(productProvider)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/products/0", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v but got %v", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		// Arrange
		productProvider := &productProviderMock{}
		mux := newHandler(productProvider)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status code %v but got %v", http.StatusNotFound, rec.Code)
		}
	})

	t.Run("ok", func(t *testing.T) {
		// Arrange
		productProvider := &productProviderMock{validSKU: 1}
		mux := newHandler(productProvider)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %v but got %v", http.StatusOK, rec.Code)
		}
		if !strings.Contains(rec.Body.String(), `"sku":1`) {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	})
}
