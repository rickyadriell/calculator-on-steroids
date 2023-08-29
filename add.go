package main

import (
	"encoding/json"
	"net/http"
)

func Add(a int, b int) int {
	return a + b
}

type Num struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Result struct {
	Value int `json:"result"`
}

func (s *Server) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req Num
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := Result{Value: Add(req.X, req.Y)}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
