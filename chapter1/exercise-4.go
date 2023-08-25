package main

import (
	"bufio"
	"fmt"
	"os"
)

func exercise4() {
  counts := make(map[string]int)
  lineFilenameMap := make(map[string][]string)

  filenames := os.Args[1:]

  if len(filenames) == 0 {
    fmt.Fprintf(os.Stderr, "dup4: no files provided")
    return
  }

  for _, filename := range filenames {
    f, err := os.Open(filename)

    if err != nil {
      fmt.Fprintf(os.Stderr, "%s\t%v\n", filename, err)
      continue
    }

    input := bufio.NewScanner(f)

    for input.Scan() {
      text := input.Text()

      // save count
      counts[text]++

      list := lineFilenameMap[text]

      fileFound := false

      for _, item := range list {
        if item == filename {
          fileFound = true
        }
      }

      if !fileFound {
        lineFilenameMap[text] = append(list, filename)
      }
    }
  }

  for line, count := range counts {
    if count > 1 {
      filenames := lineFilenameMap[line]

      fmt.Printf("%s:\t", line)

      // Print all filenames with duplicates
      for _, filename := range filenames {
        fmt.Printf("%s\t", filename)
      }

      fmt.Println()
    }
  }
}
