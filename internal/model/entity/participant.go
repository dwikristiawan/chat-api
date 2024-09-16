package entity

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	UserId uint `json:"user_id"`
	ChatId uint `json:"chat_id"`
}
