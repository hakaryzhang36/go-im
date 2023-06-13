package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendId  string
	RecvId  string
	Type    int
	Media   int
	Content string
	Pic     string
	Url     string
	Desc    string
	Amount  int
}

func (table *Message) TableName() string {
	return "message"
}
