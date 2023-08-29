package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Health(t *testing.T) {
	srv := Server{
		mux:    http.NewServeMux(),
		logger: slog.Default(),
	}
	srv.routes()
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	srv.mux.ServeHTTP(w, req)
	t.Run("check get health", func(t *testing.T) {
		want := http.StatusOK
		got := w.Result().StatusCode
		if got != want {
			t.Errorf("Server.Health() = %v != want %v", got, want)
		}
	})
	req = httptest.NewRequest("POST", "/health", nil)
	w = httptest.NewRecorder()
	srv.mux.ServeHTTP(w, req)
	t.Run("check post health", func(t *testing.T) {
		want := http.StatusMethodNotAllowed
		if got := w.Result().StatusCode; got != want {
			t.Errorf("Server.Health() = %v != want %v", got, want)
		}
	})
}
