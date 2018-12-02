package main

import (
	"fmt"
	"time"
	"sync"
)

func cacle(i int){
	time.Sleep(time.Second)
	fmt.Println(i)
	//w.Done()
}

func main(){
	var wg sync.WaitGroup
	for i :=0;i<1000;i++{
		wg.Add(1)
		go func(){
			cacle(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("all job is success!")
}

