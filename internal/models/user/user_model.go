// Package user 定义用户相关的数据模型。
package user

import "gorm.io/gorm"

// UserModel 用户数据模型，映射数据库 user 表
type UserModel struct {
	gorm.Model           // 内置 ID、CreatedAt、UpdatedAt、DeletedAt 字段
	Name   string `json:"name" gorm:"column:name"`   // 用户姓名
	Age    uint   `json:"age" gorm:"column:age"`      // 用户年龄
	Email  string `json:"email" gorm:"column:email"`  // 用户邮箱
	Gender string `json:"gender" gorm:"column:gender"` // 用户性别
}

// TableName 指定数据库表名
func (UserModel) TableName() string {
	return "user"
}
