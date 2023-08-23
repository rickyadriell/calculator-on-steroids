package main

import (
	"log/slog"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	tests := struct {
		name string
		want *Server
	}{
		name: "check new server",
		want: &Server{
			mux:    http.NewServeMux(),
			logger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
		},
	}
	t.Run(tests.name, func(t *testing.T) {
		if got := NewServer(); !reflect.DeepEqual(got, tests.want) {
			t.Errorf("NewServer() = %v, want %v", got, tests.want)
		}
	})
}
