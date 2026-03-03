package router

import "github.com/gin-gonic/gin"

// 初始化一个gin引擎
func App() *gin.Engine {
	r := gin.Default()
	return r
}
