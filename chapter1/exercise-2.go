package main

import (
  "os"
  "fmt"
)

func exercise2 () {
  for index, arg := range os.Args[1:] {
    fmt.Println(index, arg)
  }
}
