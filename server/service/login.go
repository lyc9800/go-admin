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

	// 使用 type.go 中定义的 UserWithRole 结构体
	userWithRole := &UserWithRole{}
	
	// 查询用户信息并关联角色名称
	err = models.DB.Table("sys_user").
		Select("sys_user.*, sys_role.role_name").
		Joins("left join sys_role on sys_user.role_id = sys_role.id").
		Where("sys_user.id = ?", sysUser.ID).
		First(userWithRole).Error
	
	// 如果查询角色名称失败，使用原有用户信息，并根据 roleId 设置默认角色名称
	if err != nil {
		// 使用原有的 sysUser 数据填充
		userWithRole.SysUser = *sysUser
		// 根据 roleId 设置默认角色名称
		switch sysUser.RoleId {
		case 1:
			userWithRole.RoleName = "超级管理员"
		case 2:
			userWithRole.RoleName = "管理员"
		default:
			userWithRole.RoleName = "普通用户"
		}
	}

	// 生成token
	token, err := helper.GenerateToken(sysUser.ID, sysUser.RoleId, sysUser.UserName, define.TokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	
	// 刷新token
	refreshToken, err := helper.GenerateToken(sysUser.ID, sysUser.RoleId, sysUser.UserName, define.RefreshTokenExpire)
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
		"userInfo": userWithRole,  // 返回包含 role_name 的用户信息
	})
}