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
		Handler: http.TimeoutHandler(http.DefaultServeMux, 7*time.Second, ""),
		Addr:    ":" + port,
	}

	http.HandleFunc("/backend/share", HandleShare)
	http.HandleFunc("/backend/compile", HandleCompile)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatalf("Error listening on :%v: %v", port, server.ListenAndServe())
}
