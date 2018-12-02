package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {
	seen :=make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line := input.Text()
		if !seen[line]{
			seen[line] = true
			fmt.Println(line)
		}else{
			fmt.Printf("s[%s] is exist",line)
		}
	}
	if err:=input.Err();err!=nil{
		fmt.Fprintf(os.Stderr, "dedup%v\n", err)
		os.Exit(1)
	}
}
