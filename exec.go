package main

import (
  "github.com/google/go-jsonnet"
)

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
