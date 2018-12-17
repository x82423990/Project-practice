package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = `D:\GOWORK\tmp\tail\nginx.logzap`
	config["level"] = logs.LevelInfo

	configStr, err := json.Marshal(config)
	if err !=nil{
		fmt.Println("marcshal failed ,err", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Info("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")
}
