package cmd

import (
	"vdebu.gofly.net/conf"
	"vdebu.gofly.net/global"
	"vdebu.gofly.net/router"
	"vdebu.gofly.net/utils"
)

func Start() {
	// 定义初始化模块中的哨兵错误用于错误链的追踪
	var initErr error
	// 初始化配置文件
	conf.InitConfig()
	// 初始化日志模块
	// 可以再封装一层再用这里开箱即用了
	global.Logger = conf.InitLogger()
	// 初始化数据库链接
	db, err := conf.InitDB()
	if err != nil {
		//// 向基础错误中追加新错误
		//if initErr == nil {
		//	// 判断是否初次加入
		//	initErr = err
		//} else {
		//	// 将错误包装新增 -> 错误链
		//	initErr = fmt.Errorf("%v,%w", initErr, err)
		//}

		// 更新err
		initErr = utils.AppendError(initErr, err)
	}
	global.DB = db
	// 初始化Redis
	redis, err := conf.InitRedis()
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.RDB = redis
	//global.RDB.Set("miku", "39")
	//global.RDB.Delete("miku")
	//fmt.Println(global.RDB.Get("miku"))
	//fmt.Println(global.RDB.Get("miku"))
	// 检查错误链中的错误
	if initErr != nil {
		// 先看日志组件是否初始化成功再决定是否用他来输出信息
		if global.Logger != nil {
			// 输出错误信息
			global.Logger.Error(initErr.Error())
		}
		// 直接panic
		panic(initErr.Error())
	}
	// 初始化路由模块
	router.InitRouters()
}

func Clean() {

}
