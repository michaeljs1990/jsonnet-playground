package main

import (
	"html/template"
	"net/http"
	"strings"
)

var editTemplate = template.Must(template.ParseFiles("static/index.html"))

// Example Jsonnet to show people on first visit
const defaultJsonnet = `local first = "hello";

{
  second:: "world",
  out: "%s %s!" % [first, self.second],
}`

type editData struct {
	Code string
}

func HandleEditor(w http.ResponseWriter, r *http.Request) {
	editData := editData{
		Code: defaultJsonnet,
	}

	if strings.HasPrefix(r.URL.Path, "/j/") {
		id := r.URL.Path[3:]
		code, err := store.Get(id)
		if err == nil {
			editData.Code = code
		}
	}

	// TODO: Add some error handling around
	editTemplate.Execute(w, editData)
}
