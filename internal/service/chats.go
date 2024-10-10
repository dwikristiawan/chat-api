package service

import (
	"chat-api/internal/model/dto/request/chatRequest"
	"chat-api/internal/model/dto/response/chatResponse"
	"chat-api/internal/model/entity"
	"chat-api/internal/model/entity/model"
	"chat-api/internal/repository/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type chatService struct {
	userRepository        postgres.UserRepository
	chatRepository        postgres.ChatRepository
	participantRepository postgres.ParticipantRepository
	messageRepository     postgres.MessageRepository
	broadcastService      BroadcastService
	referenceRepository   postgres.ReferenceRepository
}

type ChatService interface {
	NewGroupRoamService(context.Context, *uint, *chatRequest.NewGroupRoamRequest) error
	AddParticipantService(context.Context, *chatRequest.AddParticipantRequest) error
	NewPersonalRoamService(context.Context, *uint, *chatRequest.NewPersonalRoamRequest) error
	SentMessage(context.Context, *uint, *chatRequest.SentMessageRequest) error
	GetChatById(context.Context, *uint, *uint) error
}

func NewService(
	userRepository postgres.UserRepository,
	chatRepository postgres.ChatRepository,
	participantRepository postgres.ParticipantRepository,
	broadcastService BroadcastService,
) ChatService {
	return &chatService{
		userRepository:        userRepository,
		chatRepository:        chatRepository,
		participantRepository: participantRepository,
		broadcastService:      broadcastService,
	}
}

func (s *chatService) NewGroupRoamService(c context.Context, id *uint, req *chatRequest.NewGroupRoamRequest) error {
	var tx *gorm.DB
	newChat, err := s.chatRepository.InsertChats(c, tx, &entity.Chats{
		TypeChat: model.GroupChat,
		Name:     req.Name,
		UserId:   *id,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	maker := s.userRepository.SelectUserById(c, id)
	*req.Participant = append(*req.Participant, &maker.UUID)
	err = s.AddParticipantService(c, &chatRequest.AddParticipantRequest{
		ChatId:     newChat.ID,
		ListUserId: req.Participant,
	})
	return err
}

func (s *chatService) AddParticipantService(c context.Context, req *chatRequest.AddParticipantRequest) error {
	listUser := s.userRepository.SelectUserByUUID(c, req.ListUserId)
	var tx *gorm.DB
	var prt = new([]*entity.Participants)
	for _, user := range *listUser {
		*prt = append(*prt, &entity.Participants{
			ChatId: req.ChatId,
			UserId: user.ID,
		})
	}
	newPrt, err := s.participantRepository.InsertBatchParticipant(c, tx, prt)
	if err != nil {
		return err
	}
	roamChat, err := s.chatRepository.SelectChatWitParticipantUserById(c, &req.ChatId)
	if err != nil {
		return err
	}
	makerChatData := chatResponse.Users{
		Participant: roamChat.Maker.UUID,
		Name:        roamChat.Maker.Username,
	}
	var partcpnt []chatResponse.Users
	for _, participant := range roamChat.Participants {
		partcpnt = append(partcpnt, chatResponse.Users{
			Participant: participant.Users.UUID,
			Name:        participant.Users.Username,
		})
	}
	queMsg := make(map[string]interface{})
	for _, destination := range *newPrt {
		queMsg[fmt.Sprint(destination.Users.ID)] = chatResponse.Pack{
			Type: model.NewRoamResp,
			Data: chatResponse.NewRoamResponse{
				Name:         *roamChat.Name,
				Maker:        makerChatData,
				Participants: partcpnt,
			},
		}
	}
	s.broadcastService.BroadcastChatService(c, &queMsg)
	return nil
}
func (s *chatService) NewPersonalRoamService(c context.Context, id *uint, req *chatRequest.NewPersonalRoamRequest) error {
	var dest = s.userRepository.SelectUserByUUID(c, &[]*string{&req.Destination})
	if dest == nil || len(*dest) > 1 {
		err := errors.New(fmt.Sprintf("%s not found", req.Destination))
		log.Errorf("Error NewPersonalRoamService.destination err: %v", err)
		return err
	}

	var tx *gorm.DB
	chat, err := s.chatRepository.InsertChats(c, tx, &entity.Chats{
		TypeChat: model.PersonalChat,
		Name:     nil,
		UserId:   *id,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = s.participantRepository.InsertBatchParticipant(c, tx, &[]*entity.Participants{{
		ChatId: chat.ID,
		UserId: *id,
	}, {
		ChatId: chat.ID,
		UserId: (*(*dest)[0]).ID,
	}})
	if err != nil {
		tx.Rollback()
		return err
	}
	roamChat, err := s.chatRepository.SelectChatWitParticipantUserById(c, &chat.ID)
	if err != nil {
		return err
	}
	var prt []chatResponse.Users
	for _, participant := range roamChat.Participants {
		prt = append(prt, chatResponse.Users{
			Participant: participant.Users.UUID,
			Name:        participant.Users.Username,
		})
	}
	var queChat = make(map[string]interface{})
	queChat[fmt.Sprint(roamChat.Participants[0].Users.ID)] = chatResponse.Pack{
		Type: model.NewRoamResp,
		Data: chatResponse.NewRoamResponse{
			Name: roamChat.Participants[1].Users.Username,
			Maker: chatResponse.Users{
				Participant: chat.Maker.UUID,
				Name:        chat.Maker.Username,
			},
			Participants: prt,
		},
	}
	queChat[fmt.Sprint(roamChat.Participants[1].Users.ID)] = chatResponse.Pack{
		Type: model.NewRoamResp,
		Data: chatResponse.NewRoamResponse{
			Name: roamChat.Participants[0].Users.Username,
			Maker: chatResponse.Users{
				Participant: chat.Maker.UUID,
				Name:        chat.Maker.Username,
			},
			Participants: prt,
		},
	}
	s.broadcastService.BroadcastChatService(c, &queChat)
	err = s.SentMessage(c, id, &chatRequest.SentMessageRequest{
		ChatId:      roamChat.ID,
		MessageType: req.Content.MessageType,
		Content:     req.Content.Content,
	})
	if err != nil {
		tx.Rollback()
	}
	return err
}

func (s *chatService) SentMessage(ctx context.Context, id *uint, request *chatRequest.SentMessageRequest) error {
	var tx *gorm.DB
	chat, err := s.chatRepository.SelectChatWitParticipantUserById(ctx, &request.ChatId)
	if err != nil {
		return err
	}
	tx.Begin()
	newRef, err := s.referenceRepository.InsertReference(ctx, tx, &entity.References{ChatId: request.ChatId})
	if err != nil {
		tx.Rollback()
		return err
	}
	var sender entity.Participants
	for _, participant := range chat.Participants {
		if *id == participant.ID {
			sender = participant
		}
	}
	var newChat = new([]*entity.Messages)
	for _, participant := range chat.Participants {
		*newChat = append(*newChat, &entity.Messages{
			TypeMessage: request.MessageType,
			ChatId:      chat.ID,
			Reference:   newRef.ID,
			Content:     request.Content,
			UserId:      participant.UserId,
			SenderId:    sender.ID,
			Status:      model.StatusUnread,
		})
	}
	newChat, err = s.messageRepository.InsertBatchMessage(ctx, tx, newChat)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	msgQue := make(map[string]interface{})
	for _, messages := range *newChat {
		msgQue[fmt.Sprint(messages.UserId)] = chatResponse.Pack{
			Type: model.NewMessage,
			Data: chatResponse.MessageData{
				Id:          messages.ID,
				ChatId:      chat.ID,
				Reference:   messages.Reference,
				MessageType: messages.TypeMessage,
				Content:     messages.Content,
				Sender: chatResponse.Users{
					Participant: messages.Sender.Users.UUID,
					Name:        messages.Sender.Users.Username,
				},
				CreatedAt: messages.CreatedAt,
			},
		}
	}
	s.broadcastService.BroadcastChatService(ctx, &msgQue)
	return err
}

func (s *chatService) GetChatById(ctx context.Context, id *uint, chatId *uint) error {
	chat, err := s.chatRepository.SelectChatWitParticipantUserById(ctx, chatId)
	if err != nil {
		return err
	}
	var prt []chatResponse.Users
	for _, participant := range chat.Participants {
		if *id != participant.Users.ID {
			chat.Name = &participant.Users.Username
		}
		prt = append(prt, chatResponse.Users{
			Participant: participant.Users.UUID,
			Name:        participant.Users.Username,
		})

	}

	var que = make(map[string]interface{})
	que[fmt.Sprint(*id)] = chatResponse.Pack{
		Type: model.GetChatById,
		Data: chatResponse.NewRoamResponse{
			Name: *chat.Name,
			Maker: chatResponse.Users{
				Participant: chat.Maker.UUID,
				Name:        chat.Maker.Username,
			},
			Participants: prt,
		},
	}
	s.broadcastService.BroadcastChatService(ctx, &que)
	return err
}
