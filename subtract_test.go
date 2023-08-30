package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSubtract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{a: 2, b: 3},
			want: -1,
		},
		{
			name: "Test2",
			args: args{a: 5, b: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Subtract(t *testing.T) {
	srv := Server{
		mux:    http.NewServeMux(),
		logger: slog.Default(),
	}
	srv.routes()
	payload := strings.NewReader(`{"x":10,"y":3}`)
	req := httptest.NewRequest(http.MethodPost, "/subtract", payload)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.mux.ServeHTTP(w, req)
	t.Run("check subtract", func(t *testing.T) {
		want := http.StatusOK
		got := w.Result().StatusCode
		if got != want {
			t.Errorf("Server.Add() = %v != want %v", got, want)
		}
	})
	t.Run("check subtract", func(t *testing.T) {
		want := 7
		var result Result
		err := json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Unmarshal error")
		}
		if result.Value != want {
			t.Errorf("Server.Subtract() = %v != want %v", result.Value, want)
		}
	})
}
