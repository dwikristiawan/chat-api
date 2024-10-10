package postgres

import (
	"chat-api/internal/model/entity"
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type participantRepository struct {
	db *gorm.DB
}
type ParticipantRepository interface {
	InsertParticipant(context.Context, *entity.Participants) (*entity.Participants, error)
	InsertBatchParticipant(context.Context, *gorm.DB, *[]*entity.Participants) (*[]*entity.Participants, error)
	SelectUserParticipanByChatId(context.Context, *uint) (*[]*entity.Participants, error)
}

func NewParticipantRepository(db *gorm.DB) ParticipantRepository {
	return &participantRepository{
		db: db,
	}
}
func (r *participantRepository) InsertParticipant(ctx context.Context, participants *entity.Participants) (*entity.Participants, error) {
	err := r.db.WithContext(ctx).Create(participants).Error
	if err != nil {
		return nil, err
	}
	return participants, err
}
func (r *participantRepository) InsertBatchParticipant(ctx context.Context, tx *gorm.DB, batch *[]*entity.Participants) (*[]*entity.Participants, error) {
	err := tx.WithContext(ctx).Create(batch).Error
	if err != nil {
		log.Errorf("Error inserting participants: %v", err)
	}
	return batch, err
}
func (r *participantRepository) SelectUserParticipanByChatId(ctx context.Context, chatId *uint) (*[]*entity.Participants, error) {
	var participants = new([]*entity.Participants)
	err := r.db.WithContext(ctx).Preload("Users").First(participants, chatId).Error
	if err != nil {
		log.Errorf("Err SelectUserParticipanByChatId err: %v", err)
	}
	return participants, err
}
