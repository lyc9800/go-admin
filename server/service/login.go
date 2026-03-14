package service

import (
	"errors"
	"net/http"
	"server/define"
	"server/helper"
	"server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 用户登陆
func LoginPassword(c *gin.Context) {
	in := &LoginPassWordRequest{}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	// 根据账号密码查询结果，返回结果
	sysUser, err := models.GetUserByUsernamePassword(in.UserName, in.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
	}
	// 生成token
	token, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, define.TokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	// 刷新token
	refreshToken, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, define.RefreshTokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	data := &LoginPasswordReply{
		Token:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "登陆成功",
		"result":   data,
		"userInfo": sysUser,
	})
}
