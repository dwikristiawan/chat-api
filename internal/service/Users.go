package service

import (
	"chat-api/internal/model/dto/response/chatResponse"
	"chat-api/internal/model/dto/response/contact"
	"chat-api/internal/model/entity"
	"chat-api/internal/model/entity/model"
	"chat-api/internal/repository/postgres"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type userService struct {
	userRepository    postgres.UserRepository
	contactRepository postgres.ContactRepository
	broadcastService  BroadcastService
}
type UserService interface {
	AddContact(context.Context, *uint, *string)
}

func NewUserService(
	userRepository postgres.UserRepository,
	contactRepository postgres.ContactRepository,
) UserService {
	return &userService{
		userRepository:    userRepository,
		contactRepository: contactRepository,
	}
}
func (s *userService) AddContact(ctx context.Context, id *uint, UUID *string) {
	user := s.userRepository.SelectUserByUUID(ctx, &[]*string{UUID})
	if user == nil {
		return
	}
	var tx gorm.DB
	tx.Begin()
	newContact := s.contactRepository.InsertContact(ctx, &tx, &entity.Contacts{
		UserID:    *id,
		ContactID: (*user)[0].ID,
	})
	if newContact == nil {
		return
	}
	var que = make(map[string]interface{})
	que[fmt.Sprint(*id)] = chatResponse.Pack{
		Type: model.AddContact,
		Data: contact.Contact{
			UUID:     (*user)[0].UUID,
			Username: (*user)[0].Username,
		},
	}
	s.broadcastService.BroadcastChatService(ctx, &que)
}
