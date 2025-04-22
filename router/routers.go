package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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
	//创建ctx用于优雅退出
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
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
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", stPort),
		Handler: r,
	}
	// 异步启动服务器(之前是正常启动服务器然后异步监听退出信号)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// 启动服务器时发生了预料之外的错误
			// TODO:记录日志
			log.Printf("starting server error: %s\n", err.Error())
			return
		}
		// 启动服务器成功
	}()
	// 阻塞接受信号
	<-ctx.Done()
	// 创建ctx设置dl关闭服务器
	//ctx, cancelShutdown := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancelShutdown()
	//// 关闭服务器
	//if err := srv.Shutdown(ctx); err != nil {
	//	// TODO: 记录日志
	//	log.Printf("shutdown server error: %s\n", err.Error())
	//}

	// 创建dl用于优雅停机
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownCancel()
	// echo内置的Shutdown底层调用的还是http.Server.Shutdown
	err := r.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("graceful shutdown server error: %s", err.Error())
		return
	}
	log.Println("graceful stop server success")
}
