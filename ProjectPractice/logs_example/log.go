package main

import (
	"github.com/astaxie/beego/logs"
)

func conver(level string) int {
	switch (level) {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLoger()error{
	config:=config := make(map[string]interface{})
	config["filename"] = appConfig.logPath
}
