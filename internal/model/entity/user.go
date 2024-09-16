package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	email    string `json:"email"`
	password string `json:"password"`
}
