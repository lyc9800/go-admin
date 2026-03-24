package service

import (
	"log"
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	log.Printf("收到角色菜单 ID: %v, 长度: %d", in.MenuId, len(in.MenuId))
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
	rms := make([]*models.RoleMenu, len(in.MenuId))
	for i, _ := range rms {
		rms[i] = &models.RoleMenu{
			MenuId: in.MenuId[i],
		}
	}

	rb := &models.SysRole{
		IsAdmin: in.IsAdmin,
		Name:    in.Name,
		Remarks: in.Remarks,
		Sort:    in.Sort,
	}
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(rb).Error; err != nil {
			return err
		}
		for _, v := range rms {
			v.RoleId = rb.ID
		}
		if len(rms) > 0 {
			err = tx.Create(rms).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
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

	menuIds, err := models.GetRoleMenuId(sysRole.ID, sysRole.IsAdmin == 1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   -1,
			"msg":    "获取数据失败",
			"result": data,
		})
		return
	}
	data.MenuId = menuIds
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

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		err = models.DB.Model(&models.SysRole{}).Where("id=?", in.ID).Updates(map[string]any{
			"name":     in.Name,
			"remarks":  in.Remarks,
			"sort":     in.Sort,
			"is_admin": in.IsAdmin,
		}).Error
		if err != nil {
			return err
		}
		// 删除旧数据
		err = tx.Where("role_id=?", in.ID).Unscoped().Delete(new(models.RoleMenu)).Error
		if err != nil {
			return err
		}
		// 增加新授权的数据
		rms := make([]*models.RoleMenu, len(in.MenuId))
		for i, _ := range rms {
			rms[i] = &models.RoleMenu{
				MenuId: in.MenuId[i],
				RoleId: in.ID,
			}
		}
		if len(rms) > 0 {
			err = tx.Create(rms).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
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

// 管理员状态修改
func UpdateRoleStatus(c *gin.Context) {
	id := c.Param("id")
	is_admin := c.Param("is_admin")
	if id == "" || is_admin == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}
	// 更改状态
	err := models.DB.Model(&models.SysRole{}).Where("id=?", id).Update("is_admin", is_admin).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// 获取所有角色
func AllRole(c *gin.Context) {
	list := make([]*AllRoleListReply, 0)
	err := models.DB.Model(models.SysRole{}).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "加载数据成功",
		"result": list,
	})
}
