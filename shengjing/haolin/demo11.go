package main

import "fmt"

var container = []string{"zero","one","two"}

func main()  {
	container := map[int]string{0:"zero", 1:"one", 2:"two"}
	fmt.Printf("the container is %q.\n", container[1])
	// 断言
	_, ok := interface{}(container).(map[int]string)
	fmt.Println(ok)
}