package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建全局DB
var DB *gorm.DB

func NewGormDB() {
	// 定义数据连接信息
	dsn := "root:Mysql_2022@tcp(192.168.6.242:4306)/go_admin?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	// 自动建表
	err = db.AutoMigrate(&SysUser{})
	// 将连接返回的局部变量db赋值给全局DB变量，方便其他package引用
	DB = db
}
