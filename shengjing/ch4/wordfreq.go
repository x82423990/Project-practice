package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	wordCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		if input.Text() == "quit"{
			break
		}
		wordCount[input.Text()]++
	}
	fmt.Printf("\nWorld\tCount\n")
	for key, value := range wordCount {
		fmt.Printf("%v\t%d\n", key, value)
	}
}
