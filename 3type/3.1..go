package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

// 定义矩形
type Rect struct {
	x, y          float64
	width, height float64
}

// 定义计算面积的方法

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func main() {
	var a Integer = 1
	if a.Less(12) {
		fmt.Println(a, "LESS 2")
	}
	fmt.Println("before", a)
	a.Add(18)
	fmt.Println("affter", a)
	fmt.Println("=====================")
	rect := Rect{width: 3, height: 3}
	fmt.Println(rect.Area())
}
