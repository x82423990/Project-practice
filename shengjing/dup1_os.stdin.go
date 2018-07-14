package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		// 出现一次就加一次，设计有点巧妙.
		counts[input.Text()]++

	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)

	}
}
