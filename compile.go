package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-jsonnet"
)

func HandleCompile(w http.ResponseWriter, r *http.Request) {
	// Use MB of memory before paging out to disk
	if err := r.ParseMultipartForm(1000000); err != nil {
		fmt.Fprintf(w, "Unable to parse form: %v", err)
		return
	}

	ret, err := execJsonnetCode(r.FormValue("jsonnet"))
	if err != nil {
		fmt.Fprintf(w, "Unable to exec code with jsonnet VM: %v", err)
		return
	}

	fmt.Fprintf(w, ret)
}

func execJsonnetCode(code string) (string, error) {
	vm := jsonnet.MakeVM()
	// The first param is a filename for when an error is output. As we are doing this all in memory
	// having a filename returned for an error would likely just be confusing to a user.
	out, err := vm.EvaluateSnippet("", code)
	if err != nil {
		// Explicitly return an emptry string otherwise it will just return x as out which
		// may be confusing to people when it shows up in the editor.
		return "", err
	}

	return out, nil
}
