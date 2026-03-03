package main

import (
	"server/models"
	"server/router"
)

func main() {
	// 初始化数据库
	models.NewGormDB()
	// 初始化gin引擎
	r := router.App()
	// 监听端口
	r.Run(":9999")
}
