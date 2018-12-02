package main

import (
	"math/rand"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func insertail(tail *Student) {
	for i := 1; i < 11; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
}
func main() {
	var head Student
	//head.Name="xieyifan"
	//head.Age=19
	//head.Score=100
	insertail(&head)
	//var tail Student
	//for i := 1; i < 11; i++ {
	//	stu := Student{
	//		Name:  fmt.Sprintf("stu%d", i),
	//		Age:   rand.Intn(100),
	//		Score: rand.Float32() * 100,
	//	}
	//	tail.next = &stu
	//	tail = stu
	//}
	trans(&head)
}
