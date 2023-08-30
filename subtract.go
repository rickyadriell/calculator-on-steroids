package main

import (
	"encoding/json"
	"net/http"
)

func Subtract(a int, b int) int {
	return a - b
}

func (s *Server) Subtract() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.logger.Error("Method not allowed", "name", "Subtract handler")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req Num
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.logger.Error("Decoding", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := Result{Value: Subtract(req.X, req.Y)}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			s.logger.Error("Encoding", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
