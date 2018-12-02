package main

import "fmt"

func nonempty(s []string) []string {
	out := make([]string, 0)
	for _, k := range s {
		if k != "" {
			out = append(out, k)
		}
	}
	return out
}

func main() {
	s := []string{"a", "", "b"}
	tmp :=nonempty(s)
	fmt.Println(tmp)
	fmt.Println(s)
}
