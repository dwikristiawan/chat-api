package postgres

import (
	"chat-api/internal/model/entity"
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type referenceRepository struct {
	db *gorm.DB
}

func NewReferenceRepository(db *gorm.DB) ReferenceRepository {
	return &referenceRepository{db: db}
}

type ReferenceRepository interface {
	InsertReference(context.Context, *gorm.DB, *entity.References) (*entity.References, error)
}

func (r *referenceRepository) InsertReference(ctx context.Context, tx *gorm.DB, ref *entity.References) (*entity.References, error) {
	err := tx.WithContext(ctx).Create(*ref).Error
	if err != nil {
		log.Errorf("error insertingReference %v: %v", ref, err)
	}
	return ref, err
}
