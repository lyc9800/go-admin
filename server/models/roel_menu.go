package models

import (
	"fmt"

	"gorm.io/gorm"
)

type RoleMenu struct {
	gorm.Model
	RoleId uint `gorm:"column:role_id;type:int(11);" json:"role_id"`
	MenuId uint `gorm:"column:menu_id;type:int(11);" json:"menu_id"`
}

// 设置角色菜单表
func (table *RoleMenu) TableName() string {
	return "sys_role_menu"
}

// 获取指定角色菜单函数
func GetRoleMenuId(roleId uint, isAdmin bool) ([]uint, error) {
	tx := new(gorm.DB)
	data := make([]uint, 0)
	if isAdmin {
		tx = DB.Model(new(SysMenu)).Select("id").Order("sort ASC")
	} else {
		tx = DB.Model(new(RoleMenu)).Select("sm.id").Joins("LEFT JOIN sys_menu sm on sm.id=sys_role_menu.menu_id").
			Where("sys_role_menu.role_id=?", roleId).Order("sm.sort ASC")
	}
	err := tx.Scan(&data).Error
	return data, err
}
func GetRoleMenus(roleId uint, isAdmin bool) (*gorm.DB, error) {
	if isAdmin {
		// 管理员：查询所有菜单
		tx := DB.Model(new(SysMenu)).
			Select("id, parent_id, component_name, name, web_icon, sort, path, level").
			Order("sort ASC")
		fmt.Println("=== 管理员模式：查询所有菜单 ===")
		return tx, nil
	}

	// 普通用户：查询角色关联的菜单
	roleBasic := new(SysRole)
	err := DB.Model(roleBasic).Select("id").Where("id = ?", roleId).First(roleBasic).Error
	if err != nil {
		fmt.Printf("!!! 查询角色失败: %v !!!\n", err)
		return nil, err
	}

	fmt.Printf("=== 普通用户模式：角色ID=%d ===\n", roleBasic.ID)

	tx := DB.Model(new(RoleMenu)).
		Select("mb.id, mb.parent_id, mb.component_name, mb.name, mb.web_icon, mb.sort, mb.path, mb.level").
		Joins("LEFT JOIN sys_menu mb ON mb.id = sys_role_menu.menu_id").
		Where("sys_role_menu.role_id = ?", roleBasic.ID).
		Order("mb.sort ASC")

	return tx, nil
}
