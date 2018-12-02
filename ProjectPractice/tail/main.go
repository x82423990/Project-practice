package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
)
func main() {
	filename := "D:\\GOWORK\\tmp\\tail\\nginx.log"
	tails, er := tail.TailFile(filename, tail.Config{
		ReOpen:    false,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if er != nil {
		fmt.Println("tail file erro", er)
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg", msg.Text)
	}

}
