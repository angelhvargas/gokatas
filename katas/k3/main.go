package main

import (
	"fmt"
	"os"
)

// very simple logic to understand how to look for scripts args
//
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
