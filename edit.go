package main

import (
  "net/http"
  "fmt"
)

func (s *Server) HandleEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
