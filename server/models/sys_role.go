package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	RoleName string `gorm:"column:role_name;type:varchar(100);" json:"role_name"`      // 角色名称
	IsAdmin  int8   `gorm:"column:is_admin;type:tinyint(1);default:0" json:"is_admin"` // 是否是管理员
	Sort     int64  `gorm:"column:sort;type:int(11);default:0" json:"sort"`            // 排序
	Remarks  string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`           // 备注
}

func (table *SysRole) TableName() string {
	return "sys_role"
}

// 获取角色列表
func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(&SysRole{}).Select("id,role_name,is_admin,sort,created_at,updated_at")
	if keyword != "" {
		tx = tx.Where("name like ?", "%"+keyword+"%")
	}
	tx.Order("sort ASC")
	return tx
}

// 获取角色详情
func GetRoleDetail(id uint) (*SysRole, error) {
	role := &SysRole{}
	err := DB.Model(&SysRole{}).Where("id = ?", id).First(&role).Error
	return role, err
}
