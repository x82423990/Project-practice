package main

import (
	"fmt"
)
type test struct {
	name string
}

func (t test)say(){
	fmt.Println("hi")
}
type ReadWriter interface {
	// Read(buf []byte) (n int, err error)
	//Write(buf []byte) (n int, err error)
	say()
}

type IStream interface {
	Write(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}
var t test
var file1 ReadWriter = t
// var fil2  ReadWriter = file1

func main(){
	if _,ok := file1.(ReadWriter);ok{
		fmt.Println("实现了")
	}else {
		fmt.Println("没实现")
	}
}