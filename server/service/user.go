package service

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	in := &GetUserListRequest{NewQueryRequest()}
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	var (
		cnt  int64
		list = make([]*GetUserListReply, 0)
	)
	err = models.GetUserList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载数据成功",
		"result": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

// 新增管理员信息函数
func AddUSer(c *gin.Context) {
	in := &AddUserRequest{}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	// 验证用户名是否已存在
	var cnt int64
	err = models.DB.Model(&models.SysUser{}).Where("username=?", in.Username).Count(&cnt).Error
	if err != nil || cnt > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"code": -1,
			"msg":  "用户名已存在",
		})
		return
	}
	// 保存数据
	err = models.DB.Create(&models.SysUser{
		Email:    in.Email,
		Password: in.Password,
		Phone:    in.Phone,
		Remark:   in.Remark,
		UserName: in.Username,
	}).Error
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": -1,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

// 根据ID获取用户信息
func GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误: id不能为空",
		})
		return
	}
	uId, err := strconv.Atoi(id)
	data := &GetUserDetailReply{}
	// 获取用户信息
	sysUser, err := models.GetUserDetail(uint(uId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误: 无效的用户ID",
		})
		return
	}
	data.ID = sysUser.ID
	data.Remark = sysUser.Remark
	data.Phone = sysUser.Phone
	data.Username = sysUser.UserName
	data.Email = sysUser.Email
	// data.Password = sysUser.Password
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "获取数据成功",
		"result": data,
	})
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
	in := &UpdateUserRequest{}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}
	// 验证用户名是否已存在
	var cnt int64
	err = models.DB.Model(&models.SysUser{}).Where("username=? and id !=?", in.Username, in.ID).Count(&cnt).Error
	if err != nil || cnt > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"code": -1,
			"msg":  "用户名已存在",
		})
		return
	}
	err = models.DB.Model(new(models.SysUser)).Where("id=?", in.ID).Updates(map[string]any{
		"password": in.Password,
		"username": in.Username,
		"phone":    in.Phone,
		"email":    in.Email,
		"remark":   in.Remark,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新数据成功",
	})
}
