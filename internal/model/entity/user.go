package entity

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	UUID         string         `gorm:"type:uuid" json:"uuid"`
	Chats        []Chats        `gorm:"foreignKey:UserId"json:"chats"`
	Messages     []Messages     `gorm:"foreignKey:UserId" json:"messages"`
	Participants []Participants `gorm:"foreignKey:UserId" json:"participants"`
	Contacts     []Contacts     `gorm:"foreignKey:UserId"`
}
