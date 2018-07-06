package main

import "fmt"

func del(s []string, index int) string {
	delSemnt := s[index-1]
	if index < len(s)-1 {
		s = append(s[:index-1], s[index+1:]...)
	}
	return delSemnt
}

func main() {
	sic := []string{"a", "b", "c", "d","e","f","g"}
	fmt.Println(sic[:2])
	fmt.Println(sic[(len(sic))-1:])
	a := del(sic, 2)
	fmt.Println(sic)
	fmt.Println(a)
}
