package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-jsonnet"
)

func HandleShare(w http.ResponseWriter, r *http.Request) {
	// Use MB of memory before paging out to disk
	if err := r.ParseMultipartForm(1000000); err != nil {
		fmt.Fprintf(w, "Unable to parse form: %v", err)
		return
	}

	fmt.Fprintf(w, "share")
}
