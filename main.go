package main

import (
	"log"
	"net/http"
	"os"
)

type Server struct {
	mux *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Register() {
	s.mux.HandleFunc("/backend/compile", s.HandleCompile)

	// Catch all for serving up static assets
	s.mux.Handle("/", http.FileServer(http.Dir("./static")))
}

func main() {
	server := &Server{
		mux: http.NewServeMux(),
	}

	server.Register()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatalf("Error listening on :%v: %v", port, http.ListenAndServe(":"+port, server))
}
