package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"server/define"
	"server/models"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	err = models.GetUserList(in.Keyword).Select("id, username, email, phone, remarks, avatar, created_at, updated_at").Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
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
		Remarks:  in.Remarks,
		UserName: in.Username,
		RoleId:   in.RoleId,
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
	data.Remarks = sysUser.Remarks
	data.Phone = sysUser.Phone
	data.Username = sysUser.UserName
	data.Email = sysUser.Email
	data.RoleId = sysUser.RoleId
	data.Password = sysUser.Password
	data.RoleName = sysUser.RoleName
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
		"remarks":  in.Remarks,
		"role_id":  in.RoleId,
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

// 删除用户
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误: id不能为空",
		})
		return
	}
	uId, err := strconv.Atoi(id)
	err = models.DB.Delete(&models.SysUser{}, uId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

// 修改用户信息
func UpdateUserInfoApi(c *gin.Context) {
	// 获取用户登陆信息
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	fmt.Println(userClaim)
	in := new(UpdateUserRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}
	err = models.DB.Debug().Model(new(models.SysUser)).Where("id=?", userClaim.Id).Updates(map[string]any{
		"sex":    in.Sex,
		"avatar": in.Avatar,
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
	return
}

// 发送邮箱验证码
func SendEmail(c *gin.Context) {
	// 保存邮件验证码
	session := sessions.Default(c)
	// 安全获取 userClaim
	claimVal, exists := c.Get("userClaim")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "未获取到用户信息，请重新登录",
		})
		return
	}
	userClaim := claimVal.(*define.UserClaim)
	// userClaim := c.MustGet("userClaim").(*define.UserClaim)
	sysUser, err := models.GetUserDetail(userClaim.Id)
	toEmail := c.Query("email")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "发送邮件失败",
		})
		return
	}
	// 生成随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	content := fmt.Sprintf("【系统管理】验证码：%s，请勿泄露给其他人。", vCode)
	if toEmail == "" {
		toEmail = sysUser.Email
	}
	go utils.SendEmail(toEmail, "修改邮箱验证码", content)
	session.Set("VCode", vCode)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "保存验证码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发送成功,请查收邮件",
	})
}

// 校验验证码
func VerifyCode(c *gin.Context) {
	session := sessions.Default(c)
	VCode := session.Get("VCode")
	code := c.Query("code")
	if VCode != code {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "校验通过",
	})
}

// 更换绑定邮箱
func ChangeEmail(c *gin.Context) {
	// 1. 解析请求参数
	var req ChangeEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 2. 从session中获取用户ID
	session := sessions.Default(c)
	userClaimInterface := c.MustGet("userClaim")
	if userClaimInterface == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}

	userClaim := userClaimInterface.(*define.UserClaim)
	userId := userClaim.Id

	// 3. 查询当前用户
	var user models.SysUser
	err := models.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": -1,
				"msg":  "数据库查询失败",
			})
		}
		return
	}

	// 4. 检查新邮箱是否与当前邮箱相同
	if user.Email == req.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "新邮箱不能与当前绑定邮箱相同",
		})
		return
	}

	// 5. 检查新邮箱是否已被其他用户使用
	var count int64
	err = models.DB.Model(&models.SysUser{}).Where("email = ? AND id != ?", req.Email, userId).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "数据库查询失败",
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "该邮箱已被其他用户绑定",
		})
		return
	}

	// 6. 验证验证码（从session中获取验证码）
	sessionVCode := session.Get("VCode")
	if sessionVCode == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "验证码已过期，请重新获取",
		})
		return
	}

	if sessionVCode != req.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}

	// 7. 获取用户发送验证码的邮箱
	sentToEmail := session.Get("sentToEmail")
	targetEmail := ""
	if sentToEmail != nil {
		targetEmail = sentToEmail.(string)
	}

	// 验证邮箱是否匹配
	if targetEmail != "" && targetEmail != req.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "验证码与邮箱不匹配",
		})
		return
	}

	// 8. 保存旧邮箱
	oldEmail := user.Email

	// 9. 更新数据库
	err = models.DB.Model(&user).Update("email", req.Email).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "邮箱更新失败: " + err.Error(),
		})
		return
	}

	// 10. 清除session中的验证码
	session.Delete("VCode")
	session.Delete("sentToEmail")
	session.Save()

	// 11. 可选：发送通知邮件（异步）
	// go sendChangeEmailNotification(oldEmail, req.Email)

	// 12. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "邮箱更换成功",
		"result": gin.H{
			"old_email": oldEmail,
			"new_email": req.Email,
		},
	})
}

// 修改密码
func ChangePassword(c *gin.Context) {
	// 解析请求参数
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 安全获取用户信息
	userClaimInterface, exists := c.Get("userClaim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户未登录或登录已过期",
		})
		return
	}

	userClaim, ok := userClaimInterface.(*define.UserClaim)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "用户信息格式错误",
		})
		return
	}

	userId := userClaim.Id

	// 查询用户
	var user models.SysUser
	err := models.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": -1,
				"msg":  "数据库查询失败: " + err.Error(),
			})
		}
		return
	}

	// 验证旧密码
	if user.Password != req.OldPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "当前密码错误",
		})
		return
	}

	// 检查新旧密码是否相同
	if req.OldPassword == req.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "新密码不能与当前密码相同",
		})
		return
	}

	// 密码强度检查（可选）
	if !isPasswordStrongEnough(req.NewPassword) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "密码强度不足，建议使用字母、数字和符号的组合",
		})
		return
	}

	// 检查是否与近期使用的密码相同（如果有历史记录功能）
	// 这里可以添加检查逻辑

	// 开启事务
	tx := models.DB.Begin()

	// 更新密码
	err = tx.Model(&user).Update("password", req.NewPassword).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "密码更新失败: " + err.Error(),
		})
		return
	}

	// 保存密码修改记录（可选）
	err = savePasswordChangeHistory(tx, userId, req.NewPassword)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "保存修改记录失败: " + err.Error(),
		})
		return
	}

	// 提交事务
	tx.Commit()

	// 清除当前用户的session（强制重新登录）
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "密码修改成功，请重新登录",
		"result": ChangePasswordReply{
			Success:     true,
			Message:     "密码修改成功",
			NeedReLogin: true,
		},
	})
}

// 密码强度检查函数
func isPasswordStrongEnough(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 检查是否包含数字
	hasDigit := false
	// 检查是否包含字母
	hasLetter := false

	for _, ch := range password {
		switch {
		case ch >= '0' && ch <= '9':
			hasDigit = true
		case (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z'):
			hasLetter = true
		}
	}

	// 密码必须包含字母和数字
	return hasDigit && hasLetter
}

// 保存密码修改记录（可选功能）
func savePasswordChangeHistory(tx *gorm.DB, userId uint, newPassword string) error {
	// 这里可以添加保存密码修改历史的逻辑
	// 比如记录到 password_history 表中
	// 用于检查用户是否重复使用旧密码

	return nil
}

// 根据用户名和密码查询用户（包含角色信息）
func GetUserWithRole(username, password string) (*UserWithRole, error) {
	var result UserWithRole

	// 使用原生SQL查询
	err := DB.Raw(`
        SELECT 
            u.*,
            r.name as role_name
        FROM sys_user u
        LEFT JOIN sys_role r ON u.role_id = r.id
        WHERE u.user_name = ? AND u.password = ?
    `, username, password).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}
