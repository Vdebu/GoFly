package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// InitUserRouters 初始化用户模块的所有路由
func InitUserRouters() {
	RegRouters(func(rgPublic *echo.Group, rgAuth *echo.Group) {
		// 设置Public下的节点
		rgPublic.POST("/login", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "login succeed",
			})
		})
		// 针对当前需要进行鉴权的路由进行设置
		reAuthUser := rgAuth.Group("/user")
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
	})
}
