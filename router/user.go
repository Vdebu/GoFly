package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vdebu.gofly.net/api"
)

// InitUserRouters 初始化用户模块的所有路由
func InitUserRouters() {
	RegRouters(func(rgPublic *echo.Group, rgAuth *echo.Group) {
		// 初始化UserAPI
		userAPI := api.NewUserAPI()
		// 设置Public下的节点
		rgPublicUser := rgPublic.Group("/user")
		{
			// 层次感
			// 将对应的处理器挂载在结构体里进行调用
			rgPublicUser.POST("/login", userAPI.Login)
		}
		// 针对当前需要进行鉴权的路由进行设置
		reAuthUser := rgAuth.Group("/user")
		{
			reAuthUser.GET("", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"data": []map[string]any{
						{"id": 39, "name": "miku"},
					},
				})
			})
			reAuthUser.GET("/:id", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"data": []map[string]any{
						{"id": 39, "name": "miku"},
					},
					"inputID": c.Param("id"),
				})
			})
		}

	})
}
