package main

import "fmt"

func cacl(taskChan, reslutChan, exitChan chan int){

	for v:=range taskChan{
		flag := true
		for i:=2;i<v;i++{
			if v%i ==0{
				flag = false
				break
			}
		}
		if flag{
			reslutChan <- v
		}
	}
	exitChan <- 0
}

func main() {
	var taskChan chan int
	var reslutChan chan int
	var exitChan chan int
	taskChan = make(chan int, 100000000)
	reslutChan= make(chan int,100000000)
	exitChan= make(chan int,8)
	// push taskChan
	go func() {
		for i := 0; i < 100000; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	for i:=0;i<8;i++{
		go cacl(taskChan , reslutChan, exitChan)
	}

	//go func() {
		for i:=0;i<8;i++{
			<- exitChan
		}
		close(reslutChan)
	//}()
	for v:=range reslutChan{
		fmt.Println(v)
	}

}
