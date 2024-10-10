package chatResponse

import (
	"chat-api/internal/model/entity"
	"chat-api/internal/model/entity/model"
	"time"
)

type BatchChatResponse struct {
	ID        uint
	UpdatedAt time.Time
	Name      string
	Messages  *[]*entity.Messages
}

type ChatData struct {
	ChatId   uint   `json:"chat_id"`
	ChatName string `json:"chat_name"`
}
type MessageData struct {
	Id          uint              `json:"id"`
	ChatId      uint              `json:"chat_id"`
	Reference   uint              `json:"reference"`
	MessageType model.TypeMessage `json:"message_type"`
	Content     string            `json:"content"`
	Sender      Users             `json:"sender"`
	CreatedAt   time.Time         `json:"createdAt"`
}

type Reader struct {
	ParticipantId uint   `json:"participant_id"`
	UserId        uint   `json:"user_id"`
	ReaderId      uint   `json:"reader_id"`
	Username      string `json:"username"`
	Status        string `json:"status"`
}
