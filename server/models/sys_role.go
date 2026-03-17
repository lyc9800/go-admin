package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(100);" json:"name"`                // 角色名称
	IsAdmin int    `gorm:"column:is_admin;type:tinyint(1);default:0" json:"is_admin"` // 是否是管理员
	Sort    int64  `gorm:"column:sort;type:int(11);default:0" json:"sort"`            // 排序
	Remarks string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`           // 备注
}

func (table *SysRole) TableName() string {
	return "sys_role"
}
func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(&SysRole{}).Select("id,name,is_admin,sort,created_at,updated_at")
	if keyword != "" {
		tx = tx.Where("name like ?", "%"+keyword+"%")
	}
	tx.Order("sort ASC")
	return tx
}
