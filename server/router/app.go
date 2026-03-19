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
	// 管理员列表
	r.GET("/user", service.GetUserList)
	// 添加管理员
	r.POST("/user", service.AddUSer)
	// 获取用户详细信息
	r.GET("/user/detail/:id", service.GetUserDetail)
	// 修改用户信息
	r.PUT("/user", service.UpdateUser)
	// 删除用户
	r.DELETE("/user/:id", service.DeleteUser)
	// 角色列表
	r.GET("/role", service.GetRoleList)
	// 添加角色
	r.POST("/role", service.AddRole)
	// 获取角色详细信息
	r.GET("/role/detail/:id", service.GetRoleDetail)
	// 修改角色信息
	r.PUT("/role", service.UpdateRole)
	// 删除角色
	r.DELETE("/role/:id", service.DeleteRole)
	// 修改管理员状态
	r.PATCH("/role/:id/:is_admin", service.UpdateRoleStatus)
	// 获取菜单列表
	r.GET("/menu", service.GetMenuList)
	// 添加菜单
	r.POST("/menu", service.AddMenu)
	return r
}
