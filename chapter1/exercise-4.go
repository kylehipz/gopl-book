package main

import (
	"bufio"
	"fmt"
	"os"
)

func exercise4() {
	counts := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		saveLines("stdin", counts)
	} else {
		for _, file := range files {

      err := saveLines(file, counts)
      if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 %v\n", err)
				continue
      }
		}
	}

	for line, filenames := range counts {
		if len(filenames) > 1 {
      fmt.Printf("%s:\t", line)
      for _, filename := range filenames {
        fmt.Printf("%s ", filename)
      }
      fmt.Println()
		}
	}
}

func saveLines(filename string, counts map[string][]string) error {
  var f *os.File

  if filename == "stdin" {
    f = os.Stdin
  } else {
    file, err := os.Open(filename)

    if err != nil {
      return err
    }

    f = file
  }

	input := bufio.NewScanner(f)

	for input.Scan() {
    text := input.Text()
    list := counts[text]

    fileFound := false

    for _, existingFilename := range list {
      if existingFilename == filename {
        fileFound = true
      }
    }

    if !fileFound {
      counts[text] = append(list, filename)
    }

	}

  f.Close()

  return nil
}
