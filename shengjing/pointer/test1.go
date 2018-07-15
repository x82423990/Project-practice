package main

import "fmt"

func main()  {
	x:=1
	p:=&x
	p2 := &x
	println(p==p2)
	var z, y int
	fmt.Println(&z==&y)
}
