package chatResponse

import "chat-api/internal/model/entity/model"

type Pack struct {
	Type model.ServiceType `json:"type"`
	Data interface{}       `json:"data"`
}
type NewRoamResponse struct {
	Name         string  `json:"name"`
	Maker        Users   ` json:"maker"`
	Participants []Users ` json:"participants"`
}
type Users struct {
	Participant string `json:"participant_id"`
	Name        string `json:"name"`
}
