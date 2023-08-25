package main

import (
	"fmt"
	"os"
)

func exercise2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
