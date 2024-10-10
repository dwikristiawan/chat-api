package chatRequest

import "chat-api/internal/model/entity/model"

type NewGroupRoamRequest struct {
	Name        *string    `json:"name"`
	Participant *[]*string `json:"member"`
	Content     *Content   `json:"content"`
}
type Content struct {
	MessageType model.TypeMessage
	Content     string `json:"content"`
}

type SentMessageRequest struct {
	ChatId      uint `json:"chat_id"`
	MessageType model.TypeMessage
	Content     string `json:"content"`
}
type NewPersonalRoamRequest struct {
	Destination string   `json:"destination"`
	Content     *Content `json:"content"`
}
