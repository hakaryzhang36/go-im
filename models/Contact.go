package models

import "gorm.io/gorm"

//人员关系
type Contact struct {
	gorm.Model
	SourceId string
	TargetId string
	Type     int
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
