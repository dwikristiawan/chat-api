package chatRequest

type AddParticipantRequest struct {
	ChatId     uint       `json:"chat_id"`
	ListUserId *[]*string `json:"list_user_id"`
}
