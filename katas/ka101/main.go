package main

import (
	"bufio"
	"fmt"
	"os"
)

// unable to test on OS X. Go 1.15.x
// code should run. Use Dockerfile if in OS X

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		print(input.Text())
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
