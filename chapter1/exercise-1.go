package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise1() {
	fmt.Println(strings.Join(os.Args, " "))
}
