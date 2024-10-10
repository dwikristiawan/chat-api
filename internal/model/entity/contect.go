package entity

import (
	"gorm.io/gorm"
)

type Contacts struct {
	gorm.Model
	UserID    uint  `json:"user_id"`
	ContactID uint  `json:"contact_id"`
	Users     Users `gorm:"foreignKey:UserID"`
	Contacts  Users `gorm:"foreignKey:ContactID"`
}
