package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName  string `gorm:"column:username;type:varchar(50);not null;uniqueIndex" json:"userName"`
	Password  string `gorm:"column:password;type:varchar(36);not null" json:"-"`
	Phone     string `gorm:"column:phone;type:varchar(20);not null;uniqueIndex" json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(255);uniqueIndex" json:"wxUnionId"`
	WxOpendId string `gorm:"column:wx_opend_id;type:varchar(255);uniqueIndex" json:"wxOpendId"`
	Avatar    string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
}

// 设置表名称
func (table *SysUser) TableName() string {
	return "sys_user"
}

// 根据用户名和密码获取用户
func GetUserByUsernamePassword(username, password string) (*SysUser, error) {
	data := &SysUser{}
	err := DB.Where("username=? and password=?", username, password).First(data).Error
	return data, err
}
// 获取管理员数据列表
func GetUserList(keyword string) *gorm.DB{
	tx:=DB.Model(new(SysUser)).Select("id,username,phone,avatar,crated_at,updated_at")
	if keyword !="" {
		tx.Where("username LIKE ?","%"+keyword+"%")
	}
	return tx
}
