package main

import (
  "fmt"
  "os"
  "strings"
)

func dup3() {
  counts := make(map[string]int)

  files := os.Args[1:]

  for _, filename := range files {
    data, err := os.ReadFile(filename)

    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
    }

    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
    }
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
