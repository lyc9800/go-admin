package router

import (
	"server/define"
	"server/middleware"
	"server/service"

	"github.com/gin-gonic/gin"
)

// 初始化一个gin引擎
func App() *gin.Engine {
	r := gin.Default()
	// 添加跨域中间件
	r.Use(middleware.Cors())
	// 静态文件路由
	r.Static("/uploadFile", define.StaticResource)
	// 根据用户名和密码登录的路由
	r.POST("/login/password", service.LoginPassword)
	// 登录认证
	loginAuth := r.Group("/").Use(middleware.LoginAuthCheck())
	// 上传图片文件
	loginAuth.POST("/upload/file", service.UploadFile)
	/* 用户管理 ----start ---- */
	// 管理员列表
	loginAuth.GET("/user", service.GetUserList)
	// 添加管理员
	loginAuth.POST("/user", service.AddUSer)
	// 获取用户详细信息
	loginAuth.GET("/user/detail/:id", service.GetUserDetail)
	// 修改用户信息
	loginAuth.PUT("/user", service.UpdateUser)
	// 更新用户信息
	loginAuth.PUT("/user/info", service.UpdateUserInfoApi)
	// 发送邮件
	loginAuth.GET("/user/sendemail", service.SendEmail)
	// 删除用户
	loginAuth.DELETE("/user/:id", service.DeleteUser)
	/* 用户管理 ----end ---- */
	/* 角色管理 ----start ---- */
	// 获取所有角色
	loginAuth.GET("/role/all", service.AllRole)
	// 角色列表
	loginAuth.GET("/role", service.GetRoleList)
	// 添加角色
	loginAuth.POST("/role", service.AddRole)
	// 获取角色详细信息
	loginAuth.GET("/role/detail/:id", service.GetRoleDetail)
	// 修改角色信息
	loginAuth.PUT("/role", service.UpdateRole)
	// 删除角色
	loginAuth.DELETE("/role/:id", service.DeleteRole)
	// 修改管理员状态
	loginAuth.PATCH("/role/:id/:is_admin", service.UpdateRoleStatus)
	/* 角色管理 ----end ---- */
	/* 菜单管理 ----start ---- */
	// 获取菜单列表
	loginAuth.GET("/menu", service.GetMenuList)
	// 添加菜单
	loginAuth.POST("/menu", service.AddMenu)
	// 更新菜单
	loginAuth.PUT("/menu", service.UpdateMenu)
	// 删除菜单
	loginAuth.DELETE("/menu/:id", service.DelMenuApi)
	/* 菜单管理 ----end ---- */
	return r
}
