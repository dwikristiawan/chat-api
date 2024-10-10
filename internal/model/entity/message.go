package entity

import (
	"chat-api/internal/model/entity/model"
	"gorm.io/gorm"
)

type Messages struct {
	gorm.Model
	TypeMessage model.TypeMessage `json:"type_message"`
	ChatId      uint              `json:"chat_id"`
	Reference   uint              `json:"reference_id"`
	Content     string            `json:"content"`
	UserId      uint              `json:"user_id"`
	SenderId    uint              `json:"sender_id"`
	Sender      Participants      `json:"sender" gorm:"foreignKey:ParticipantId"`
	Status      model.ReadStatus  `json:"status"`
}
