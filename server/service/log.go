package service

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)
func GetLogList(c *gin.Context) {
	// 初始化请求结构体
	in := &GetLogListRequest{
		QueryRequest: NewQueryRequest(), // 注意这里需要解引用，因为 NewQueryRequest 返回的是 *QueryRequest
	}
	
	// 绑定查询参数
	if err := c.ShouldBindQuery(in); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常: " + err.Error(),
		})
		return // 必须 return，否则会继续执行
	}
	
	var (
		cnt  int64
		list = make([]*GetLogListReply, 0)
	)
	
	// 构建查询
	db := models.GetLogList(in.Keyword) // 注意是大写的 Keyword
	
	// 执行查询
	err := db.Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查询失败: " + err.Error(),
		})
		return
	}
	
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"result": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}