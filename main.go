package main

import (
  "net/http"
  "log"
  "os"
)

type Server struct {
  mux *http.ServeMux
}

func (s *Server) Init() {
  s.mux.HandleFunc("/backend", s.HandleEdit)

  // Catch all for serving up static assets
  s.mux.Handle("/", http.FileServer(http.Dir("./static")))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func main() {
  server := &Server{
    mux: http.NewServeMux(),
  }

  server.Init()

  port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
  }

  log.Fatalf("Error listening on :%v: %v", port, http.ListenAndServe(":"+port, server))
}
