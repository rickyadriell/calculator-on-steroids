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
		if r.Method == http.MethodPost {
			var num Num
			err := json.NewDecoder(r.Body).Decode(&num)
			if err != nil {
				s.logger.Error("Unmarshall error", "error", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			sum := Add(num.X, num.Y)
			s.logger.Info("Sum", "a=", num.X, "b=", num.Y, "Result", sum)
			result := Result{Value: sum}
			err = json.NewEncoder(w).Encode(result)
			if err != nil {
				s.logger.Error("Marshall error", "error", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			s.logger.Info("Status OK", "name", "add")
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
