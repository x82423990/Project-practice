package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Shoud be non-negative number!")
		return
	}
	return a + b, nil
}

// 不定参数

func function1(args ...int) {
	for i, j := range args {
		fmt.Printf("%v=%v\n", i, j)
	}
}

// 传入不定长， 不定类型的
func myPrintf(args ...interface{}) {
	for _, i := range args {
		switch i.(type) {
		case int:
			fmt.Println(i, "is int value")
		case string:
			fmt.Println(i, "is string")
		case int64:
			fmt.Println(i, "is int64 value")
		case float32:
			fmt.Println(i, "is int32 value")
		case float64:
			fmt.Println(i, "is int64 value")
		default:
			fmt.Println(i, "is unkown type")
		}
	}
}
func main() {
	ret, err := add(1, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 省略号可以打散slice
	function1(a...)
	// 不定长，不定类型
	var v1  = 1
	var v2 int64 = 234
	var v3  = "hello"
	var v4  = 1.234
	myPrintf(v1, v2, v3, v4)

}
