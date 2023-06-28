package models

import "gorm.io/gorm"

// 关系表
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系
	TargetId uint //对应的谁
	Type     int  // 类型 0 1 2
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
