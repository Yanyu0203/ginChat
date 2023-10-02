package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   uint   //sender
	TargetId uint   //receiver
	Type     string //群聊...
	Media    int    //文字、图片等
	Content  string
	Pic      string
	Url      string
	Desc     string
	Amount   int //数字统计
}

func (table *Message) TableName() string {
	return "message"
}
