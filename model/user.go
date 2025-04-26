package model

import "gorm.io/gorm"

type User struct {
	// 嵌入gorm默认字段
	gorm.Model
	NickName string `gorm:"size:64;not null"` // 长度64个字符不为空
	UserName string `gorm:"size:128"`
	Avatar   string `gorm:"size:255"`
	Number   string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:128;not null"`
}
