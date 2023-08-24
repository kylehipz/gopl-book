package main

import (
  "os"
  "fmt"
  "strings"
)

func echo3 () {
  fmt.Println(strings.Join(os.Args[1:], " "))
}
