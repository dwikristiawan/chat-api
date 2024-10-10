package postgres

import (
	"chat-api/internal/model/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}
type MessageRepository interface {
	InsertMessage(context.Context, *gorm.DB, *entity.Messages) (*entity.Messages, error)
	InsertBatchMessage(context.Context, *gorm.DB, *[]*entity.Messages) (*[]*entity.Messages, error)
	SelectByChatIdWithPage(context.Context, *uint, *int, *int) (*[]*entity.Messages, error)
	SelectMessagePackById(context.Context, *uint) *entity.Users
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) InsertMessage(ctx context.Context, tx *gorm.DB, messages *entity.Messages) (*entity.Messages, error) {
	err := tx.WithContext(ctx).Create(messages).Error
	return messages, err
}
func (r *messageRepository) InsertBatchMessage(ctx context.Context, tx *gorm.DB, messages *[]*entity.Messages) (*[]*entity.Messages, error) {
	err := tx.WithContext(ctx).Create(&messages).Error
	if err != nil {
		log.Errorf("InsertBatchMessage() error = %v", err)
	}
	return messages, err
}
func (r *messageRepository) SelectByChatIdWithPage(ctx context.Context, chatId *uint, limit *int, offset *int) (*[]*entity.Messages, error) {
	var messages []*entity.Messages
	err := r.db.WithContext(ctx).Select("*").Limit(*limit).Offset(*offset).Find(&messages).Where("chat_id", chatId).Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}
	return &messages, nil
}
func (r *messageRepository) SelectMessagePackById(c context.Context, msgId *uint) *entity.Users {
	var msg = new(entity.Users)
	err := r.db.WithContext(c).
		Preload("Chats").
		Preload("Chats.Messages", "id", *msgId).
		Preload("Chats.Messages.Readers").
		Preload("Chats.Participants").
		Find(msg).Error
	if err != nil {
		log.Errorf("Err SelectChatByMessageId.r.db.WithContext(ctx).Find(msg) err: %v", err)
		return nil
	}
	a, _ := json.Marshal(msg)
	fmt.Println("=================================================")
	fmt.Println(string(a))
	fmt.Println("=================================================")
	return msg
}
