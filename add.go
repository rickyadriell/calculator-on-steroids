package main

import (
	"encoding/json"
	"net/http"
)

func Add(a int, b int) int {
	return a + b
}

type Num struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	Num int `json:"num"`
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
			sum := Add(num.A, num.B)
			s.logger.Info("Sum", "a=", num.A, "b=", num.B, "Result", sum)
			result := Result{Num: sum}
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
