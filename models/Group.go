package models

import "gorm.io/gorm"

//群信息
type Group struct {
	gorm.Model
	GroupName string
	OwnerId   uint
	Icon      string
	Type      int
	Desc      string
}

func (table *Group) TableName() string {
	return "group"
}
