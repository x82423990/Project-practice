package main

import (
	"os"
	"bufio"
	"fmt"
	"unicode"
)

func main() {
	count :=0
	invaild := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		fmt.Println(r, n, err)
		if r == unicode.ReplacementChar && n == 1 {
			invaild++
			continue
		}
		count++
		fmt.Println("222222222222222",count)
	}

}
