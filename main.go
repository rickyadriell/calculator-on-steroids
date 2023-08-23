package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type Server struct {
	mux    *http.ServeMux
	logger *slog.Logger
}

func NewServer() *Server {
	return &Server{
		mux:    http.NewServeMux(),
		logger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
	}
}

func run() error {
	s := NewServer()

	s.routes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}
	s.logger.Info("Serving ...", "port", server.Addr)
	return server.ListenAndServe()
}
