package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/backend/compile", HandleCompile)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatalf("Error listening on :%v: %v", port, server.ListenAndServe())
}
