package models

import "gorm.io/gorm"

//聊天群表
type Group struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

func (table *Group) TableName() string {
	return "group"
}
