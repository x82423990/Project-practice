package main

import "fmt"

func main() {
	em := [10]int{78, 3, 4, 56, 7, 6, 7, 88, 10,99}
	fmt.Println(em)
	for i:=0;i<len(em);i++{
		for j:=0;j<j-i;j++{
			if em[j]>em[j+1]{
				em[j], em[j+1] =em[j+1], em[j]
			}
			break
		}
	}
	fmt.Println(em)
}
