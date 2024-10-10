package entity

import "gorm.io/gorm"

type References struct {
	gorm.Model
	ChatId   uint       `json:"chat_id"`
	Chats    Chats      `gorm:"foreignKey:ChatId" json:"chats"`
	Messages []Messages `gorm:"foreignKey:Reference" json:"messages"`
}
