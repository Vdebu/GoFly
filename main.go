package main

import "vdebu.gofly.net/cmd"

func main() {
	// 启动服务器
	cmd.Start()
	// 回收服务器资源
	defer cmd.Clean()
}
