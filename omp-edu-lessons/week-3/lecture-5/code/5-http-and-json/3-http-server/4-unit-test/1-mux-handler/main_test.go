package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_helloHandler(t *testing.T) {
	t.Run("wrong method", func(t *testing.T) {
		// Arrange
		mux := newRouter()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/hello", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status code %v but got %v", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("ok", func(t *testing.T) {
		// Arrange
		mux := newRouter()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %v but got %v", http.StatusOK, rec.Code)
		}
		if rec.Body.String() != "Hello World" {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	})
}
