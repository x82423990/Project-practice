package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	a := f(1, 2)
	fmt.Println(a)
	var j int = 5
	b := func()(func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	b()
	j *=2
	b()
}
