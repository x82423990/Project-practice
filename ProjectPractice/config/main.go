package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
)

func main() {

	conf, err := config.NewConfig("ini", `D:\GOWORK\tmp\test.conf`)
	if err != nil {
		fmt.Println("filed load config err", err)
	}
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server port failed,err", err)
	}
	log_level := conf.String("logs::log_level")
	if len(log_level) == 0 {
		fmt.Println("read config log_level  failed,err", err)
		log_level = "Debug"
	}
	fmt.Printf("port:%d\nlevel:%s", port, log_level)

}
