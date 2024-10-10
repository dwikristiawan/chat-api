package postgres

import (
	"chat-api/internal/model/entity"
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}
type ChatRepository interface {
	InsertChats(context.Context, *gorm.DB, *entity.Chats) (*entity.Chats, error)
	//SelectByParticipantIdWithPage(context.Context, *uint, *int, *int) (*[]*entity.Chats, error)
	//SelectPersonalChats(context.Context, *uint, *uint) *entity.Chats
	SelectChatWitParticipantUserById(context.Context, *uint) (*entity.Chats, error)
	//SelectMessageUnreadByUserId(context.Context, *uint) *[]*entity.Chats
	//SelectChatByMessageId(context.Context, *uint) *entity.Chats
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}
func (r *chatRepository) InsertChats(c context.Context, tx *gorm.DB, chats *entity.Chats) (*entity.Chats, error) {
	err := tx.WithContext(c).Create(chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepository) SelectChatWitParticipantUserById(c context.Context, id *uint) (*entity.Chats, error) {
	var chats *entity.Chats
	err := r.db.WithContext(c).
		Preload("Participant").
		Preload("Participant.User").
		Find(&chats, *id).Error
	if err != nil {
		log.Error(err.Error())
	}
	return chats, err
}
