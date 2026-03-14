package router

import (
	"server/middleware"
	"server/service"

	"github.com/gin-gonic/gin"
)

// 初始化一个gin引擎
func App() *gin.Engine {
	r := gin.Default()
	// 添加跨域中间件
	r.Use(middleware.Cors())
	// 根据用户名和密码登录的路由
	r.POST("/login/password", service.LoginPassword)
	// 管理员列表 start
	r.GET("/user",service.GetUserList)
	// 管理员列表 end
	return r
}
