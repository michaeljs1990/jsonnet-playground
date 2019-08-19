package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
)

func stringToHash(code string) string {
	hash := sha256.Sum256([]byte(code))
	b := base64.URLEncoding.EncodeToString(hash[0:])
	// Web sites don’t always linkify a trailing underscore, making it seem like
	// the link is broken. If there is an underscore at the end of the substring,
	// extend it until there is not.
	hashLen := 11
	for hashLen <= len(b) && b[hashLen-1] == '_' {
		hashLen++
	}
	return string(b)[:hashLen]
}

func HandleShare(w http.ResponseWriter, r *http.Request) {
	// Use MB of memory before paging out to disk
	if err := r.ParseMultipartForm(1000000); err != nil {
		fmt.Fprintf(w, "Unable to parse form: %v", err)
		return
	}

	code := r.FormValue("jsonnet")
	id := stringToHash(code)
	if err := store.Store(id, code); err != nil {
		fmt.Fprintf(w, "Unable to store code: %v", err)
	}

	fmt.Fprintf(w, id)
}
