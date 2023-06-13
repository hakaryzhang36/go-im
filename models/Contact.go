package models

import (
	"gorm.io/gorm"
	"zhangteam.org/im-project/utils"
)

// 人员关系
type Contact struct {
	gorm.Model
	SourceId uint
	TargetId uint
	Type     int
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []User {
	contacts := []Contact{}
	ids := []uint64{}
	utils.DB.Where(&Contact{SourceId: userId, Type: 1}).Find(&contacts)
	for _, v := range contacts {
		ids = append(ids, uint64(v.TargetId))
	}
	users := []User{}
	utils.DB.Where("id in ?", ids).Find(&users)
	return users
}
