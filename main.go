package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var (
	store PersistJsonnet
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Add in toggle for the different types of backends
	store = InMemory{
		store: map[string]string{},
	}

	server := &http.Server{
		Handler: http.TimeoutHandler(http.DefaultServeMux, 7*time.Second, ""),
		Addr:    ":" + port,
	}

	http.HandleFunc("/", HandleEditor)
	http.HandleFunc("/backend/share", HandleShare)
	http.HandleFunc("/backend/compile", HandleCompile)

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatalf("Error listening on :%v: %v", port, server.ListenAndServe())
}
