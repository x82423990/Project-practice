package main

import (
	"context"
	"time"
	"net/http"
	"fmt"
)

func main() {

}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c :=make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err !=nil{
		fmt.Println("http request failed, err", err)
		return
	}
	go func() {
		resp, err := client.
	}()
}

type Result struct {
	r *http.Response
	err error
}