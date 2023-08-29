package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
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
			args: args{a: 1, b: 2},
			want: 3,
		},
		{
			name: "Test2",
			args: args{a: 2, b: 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Add(t *testing.T) {
	srv := Server{
		mux:    http.NewServeMux(),
		logger: slog.Default(),
	}
	srv.routes()
	input := Num{X: 10, Y: 3}
	inputJSON, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewReader(inputJSON))
	w := httptest.NewRecorder()
	srv.mux.ServeHTTP(w, req)
	t.Run("check add", func(t *testing.T) {
		want := http.StatusOK
		got := w.Result().StatusCode
		if got != want {
			t.Errorf("Server.Add() = %v != want %v", got, want)
		}
	})
	t.Run("check add", func(t *testing.T) {
		want := 13
		var result Result
		err := json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Unmarshal error")
		}
		if result.Value != want {
			t.Errorf("Server.Add() = %v != want %v", result.Value, want)
		}
	})
}
