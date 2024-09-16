package entity

import "gorm.io/gorm"

type Readers struct {
	gorm.Model
	UserId uint `json:"user_id"`
}
