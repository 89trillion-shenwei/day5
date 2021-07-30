package main

import (
	"day5/app/http/httpserver"
	"day5/logUtil"
)

func main() {
	//初始化日志
	logUtil.Init()
	//初始化服务器
	httpserver.Init()
}
