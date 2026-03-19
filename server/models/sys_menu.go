package models

import "gorm.io/gorm"

// 获取菜单
type SysMenu struct {
	gorm.Model
	ParentId      uint   `gorm:"column:parent_id;type:int(11)" json:"parent_id"`                // 父级ID
	Name          string `gorm:"column:name;type:varchar(100)" json:"name"`                     // 菜单名称
	WebIcon       string `gorm:"column:web_icon;type:varchar(100)" json:"web_icon"`             // 图标
	Path          string `gorm:"column:path;type:varchar(255)" json:"path"`                     // 菜单路径
	Sort          int    `gorm:"column:sort;type:int(11);default:0" json:"sort"`                // 排序规则
	Level         uint   `gorm:"column:level;type:tinyint(1);default:0" json:"level"`           // 0表示目录，1 表示菜单 ,2 表示按钮
	ComponentName string `gorm:"column:component_name;type:varchar(100)" json:"component_name"` // 组件名称
}

// 定义建表语句
func (table *SysMenu) TableName() string {
	return "sys_menu"
}

// 获取菜单列表
func GetMenuList() *gorm.DB {
	return DB.Model(&SysMenu{}).Select("id,name,web_icon,path,sort,level,component_name,parent_id").Order("sort ASC")
}
