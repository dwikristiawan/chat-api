package entity

import "gorm.io/gorm"

type Messages struct {
	gorm.Model
	ChatId  uint   `json:"chat_id"`
	UserId  uint   `json:"user_id"`
	Content string `json:"content"`
}
