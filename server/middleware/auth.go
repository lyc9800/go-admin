package middleware

import (
	"fmt"
	"net/http"
	"server/helper"
	"server/models"

	"github.com/gin-gonic/gin"
)

func LoginAuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaim, err := helper.AnalyzeToken(c.GetHeader("AccessToken"))
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 60403,
				"msg":  "请先登录",
			})
			return // 修复：添加 return
		}
		fmt.Println(userClaim)
		// 检查角色ID是否有效
		if userClaim.RoleId == 0 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 60404,
				"msg":  "无权限",
			})
			return // 修复：添加 return
		}

		// 查询角色信息
		sysRole := new(models.SysRole)
		err = models.DB.Model(new(models.SysRole)).
			Select("is_admin").
			Where("id = ?", userClaim.RoleId).
			First(sysRole).Error // 修复：添加 First() 并将结果存入 sysRole

		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 60405,
				"msg":  "无此角色",
			})
			return // 修复：添加 return
		}

		// 设置管理员标志
		if sysRole.IsAdmin == 1 {
			userClaim.IsAdmin = true
		} else {
			userClaim.IsAdmin = false
		}

		c.Set("UserClaim", userClaim)
		c.Next()
	}
}
