package main

import "fmt"

func my(num int) int{
	if num<3{
		return 1
	}
	tmp := my(num-1) + my(num-2)
	return tmp
}

func test(a string) string{
	return a
}


func main() {
	fmt.Println(my(3))

}