package main

import (
  "net/http"
  "log"
  "os"
)

type Server struct {
  mux *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func main() {
  server := &Server{
    mux: http.NewServeMux(),
  }

  port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
  }

  log.Fatalf("Error listening on :%v: %v", port, http.ListenAndServe(":"+port, server))
}
