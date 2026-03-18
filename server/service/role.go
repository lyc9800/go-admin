package service

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取角色函数
func GetRoleList(c *gin.Context) {
	in := &GetRoleListRequest{NewQueryRequest()}
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	var (
		cnt  int64
		list = make([]*GetRoleListReply, 0)
	)
	err = models.GetRoleList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
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
			"total": cnt,
			"list":  list,
		},
	})
}

// 添加角色
func AddRole(c *gin.Context) {
	in := &AddRoleRequest{}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	var cnt int64
	err = models.DB.Model(&models.SysRole{}).Where("name = ?", in.Name).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "添加角色失败",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "角色名称已存在",
		})
		return
	}
	err = models.DB.Create(&models.SysRole{
		Name:    in.Name,
		Remarks: in.Remarks,
		Sort:    in.Sort,
		IsAdmin: in.IsAdmin,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "添加角色失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加角色成功",
	})
}

// 获取角色ID信息
func GetRoleDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "错误: ID不能为空",
		})
		return
	}
	uID, err := strconv.Atoi(id)
	data := &GetRoleDetailReply{}
	sysRole, err := models.GetRoleDetail(uint(uID))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "角色不存在",
		})
		return
	}
	data.ID = sysRole.ID
	data.Name = sysRole.Name
	data.Remarks = sysRole.Remarks
	data.Sort = sysRole.Sort
	data.IsAdmin = sysRole.IsAdmin
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "获取角色信息成功",
		"result": data,
	})
}

// 修改角色
func UpdateRole(c *gin.Context) {
	in := &UpdateRoleRequest{}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	var cnt int64
	err = models.DB.Model(&models.SysRole{}).Where("id != ? and name=?", in.ID, in.Name).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "修改角色失败",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "角色名称已存在",
		})
		return
	}
	err = models.DB.Model(&models.SysRole{}).Where("id=?", in.ID).Updates(map[string]any{
		"name":     in.Name,
		"remarks":  in.Remarks,
		"sort":     in.Sort,
		"is_admin": in.IsAdmin,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "修改角色失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改角色成功",
	})
}

// 删除角色
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误: id不能为空",
		})
		return
	}
	uId, err := strconv.Atoi(id)
	err = models.DB.Delete(&models.SysRole{}, uId).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除数据成功",
	})
}
