package main

import (
	"fmt"
)

func send(ch chan int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a

}
func recv(ch chan int, exitChan chan struct{}) {
	for {
		i, ok := <-ch
		if !ok{
			fmt.Println("ok",ok)
			break
		}
		fmt.Println(i)
	}
	var a struct{}
	exitChan <- a
	close(exitChan)
}

func main() {
	var ch chan int
	var exitCh chan struct{}
	ch = make(chan int, 10)
	exitCh = make(chan struct{}, 1)
	go send(ch,exitCh)
	go recv(ch, exitCh)
	for _ = range exitCh {

	}
}
