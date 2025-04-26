package cmd

import (
	"vdebu.gofly.net/conf"
	"vdebu.gofly.net/global"
	"vdebu.gofly.net/router"
)

func Start() {

	// 初始化配置文件
	conf.InitConfig()
	// 初始化日志模块
	// 可以再封装一层再用这里开箱即用了
	global.Logger = conf.InitLogger()
	// 初始化路由模块
	router.InitRouters()
}

func Clean() {

}
