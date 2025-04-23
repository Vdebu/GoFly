package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAPI struct {
}

func NewUserAPI() UserAPI {
	return UserAPI{}
}

// true表示是必须的

// @tags 用户管理
// @summary 用户登录
// @description 用户登录详情描述
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @success 200 {object} string "登录成功"
// @failure 401 {object} string "登入失败"
// @router       /api/v1/public/user/login [post]
func (m UserAPI) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
