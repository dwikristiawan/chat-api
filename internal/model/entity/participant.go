package entity

import "gorm.io/gorm"

type Participants struct {
	gorm.Model
	ChatId uint  `json:"chat_id"`
	UserId uint  `json:"user_id"`
	Users  Users `gorm:"foreignKey:UserId" json:"user"`
}
