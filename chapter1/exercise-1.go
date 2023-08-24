package main

import (
  "os"
  "fmt"
  "strings"
)

func exercise1 () {
  fmt.Println(strings.Join(os.Args, " "))
}
