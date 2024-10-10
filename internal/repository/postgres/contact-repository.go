package postgres

import (
	"chat-api/internal/model/entity"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

type ContactRepository interface {
	InsertContact(context.Context, *gorm.DB, *entity.Contacts) *entity.Contacts
	SelectContactByUserId(context.Context, *uint, *int, *int) (*[]*entity.Contacts, error)
}

func (r *contactRepository) InsertContact(ctx context.Context, tx *gorm.DB, contact *entity.Contacts) *entity.Contacts {
	err := tx.WithContext(ctx).Create(contact).Error
	if err != nil {
		log.Errorf("Error inserting contact: %v", err)
		return nil
	}
	return contact
}
func (r *contactRepository) SelectContactByUserId(ctx context.Context, userId *uint, limit *int, offset *int) (*[]*entity.Contacts, error) {
	var contacts []*entity.Contacts
	if err := r.db.WithContext(ctx).Preload("Users", "ID = ?", *userId).Limit(*limit).Offset(*offset).Find(&contacts).Error; err != nil {
		log.Errorf("Error selecting contacts: %v", err)
	}
	return &contacts, nil
}
