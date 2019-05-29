package main

import (
  "net/http"
  // "github.com/google/go-jsonnet"
)

func main() {
  out , err := execJsonnetCode("{ x: 5, y: self.x,, }")
  if err != nil {
    println(err.Error())
  }

  println(out)
}
