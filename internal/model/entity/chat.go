package entity

import "gorm.io/gorm"

type chats struct {
	gorm.Model
	Name string `json:"name"`
}
