package main

import (
	"Project/conf"
	"Project/server"
	"Project/service"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	go service.Listen()
	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
