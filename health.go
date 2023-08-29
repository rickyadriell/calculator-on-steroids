package main

import "net/http"

func (s *Server) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			s.logger.Info("Bad request", "name", "get health")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		s.logger.Info("Status Ok", "name", "get health")
		w.WriteHeader(http.StatusOK)
	}
}
