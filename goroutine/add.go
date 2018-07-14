package main

import (
	"fmt"
)

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	//for i:=0;i<10;i++{
	//	go Add(i,i)
	//}
	//chs := make([]chan int, 10)
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go Count(chs[i])
	//}
	//for _,ch := range chs{
	//	<-ch
	//}

	// ch := make(chan int, 1)
	//var ch chan  int
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}

}
