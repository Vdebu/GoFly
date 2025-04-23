package main

import (
	"vdebu.gofly.net/cmd"
	"vdebu.gofly.net/router"
)

// @title go-web develop
// @version 1.0.0
// @description new start
func main() {
	// 初始化服务器配置信息
	cmd.Start()
	// 初始化路由模块
	router.InitRouters()
	// 回收服务器资源
	defer cmd.Clean()
}
