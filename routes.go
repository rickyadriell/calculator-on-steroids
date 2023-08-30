package main

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.Health())
	s.mux.HandleFunc("/add", s.Add())
	s.mux.HandleFunc("/subtract", s.Subtract())
}
