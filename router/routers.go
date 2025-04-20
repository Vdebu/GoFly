package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// IFnRegRoute IFnRegisterRoute 存储用于鉴权的路由组与公开的路由组
type IFnRegRoute = func(rgPublic *echo.Group, rgAuth *echo.Group)

// 根据配置文件中的信息决定路由模块是否启用
var (
	// echo function Routers
	// 存储所有的路由方法
	efnRouters []IFnRegRoute
)

// RegRouters 初始化所有路由函数
func RegRouters(fn IFnRegRoute) {
	if fn == nil {
		return
	}
	// 若不为空则添加到切片中
	efnRouters = append(efnRouters, fn)
}

// 初始化基础模块的路由
func regBaseRouters() {
	// 注入用户模块
	InitUserRouters()
}

// InitRouters InitRouter 初始化所有路由
func InitRouters() {
	r := echo.New()
	// 定义api公开访问路由
	rgPublic := r.Group("/api/v1/public")
	// 定义权限敏感的路由
	rgAuth := r.Group("/api/v1")
	// 初始化所有基础路由(user...)
	regBaseRouters()
	// 注册所有的模块
	for _, fnRegRoute := range efnRouters {
		// 将定义的默认参数传入作为API平台基础初始化其余路由
		fnRegRoute(rgPublic, rgAuth)
	}
	// 初始化服务器其他的相关信息
	stPort := viper.GetInt("server.port")
	// 如果viper没有读取到预期的整数数据则默认为0
	if stPort == 0 {
		// 设置默认值
		stPort = 3939
	}
	err := r.Start(fmt.Sprintf(":%d", stPort))
	if err != nil {
		panic(fmt.Sprintf("starting server error: %s", err.Error()))
	}

}
