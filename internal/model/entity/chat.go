package entity

import (
	"chat-api/internal/model/entity/model"
	"gorm.io/gorm"
)

type Chats struct {
	gorm.Model
	TypeChat     model.TypeChat `json:"typeChat"`
	Name         *string        `json:"name"`
	UserId       uint           `json:"userId"`
	Maker        Users          `gorm:"foreignKey:UserId" json:"created"`
	Participants []Participants `gorm:"foreignKey:ChatId" json:"participants"`
	Reference    []References   `gorm:"foreignKey:ChatId" json:"Reference"`
}
