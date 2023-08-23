package main

import "net/http"

func (s *Server) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.logger.Info("Status Ok", "name", "get health")
			w.WriteHeader(http.StatusOK)
			return
		} else {
			s.logger.Info("Bad request", "name", "get health")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
