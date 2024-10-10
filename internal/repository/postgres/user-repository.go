package postgres

import (
	"chat-api/internal/model/entity"
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}
type UserRepository interface {
	SelectUserById(context.Context, *uint) *entity.Users
	SelectUserByUUID(context.Context, *[]*string) *[]*entity.Users
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) SelectUserById(c context.Context, id *uint) *entity.Users {
	var user entity.Users
	err := r.db.WithContext(c).First(&user, "id = ?", id).Error
	if err != nil {
		log.Errorf("userRepository SelectUserById error, %v", err)
	}
	return &user
}
func (r *userRepository) SelectUserByUUID(c context.Context, uuid *[]*string) *[]*entity.Users {
	var users = new([]*entity.Users)
	if err := r.db.WithContext(c).Where("uuid IN ", *uuid).Find(users); err != nil {
		log.Errorf("userRepository SelectUserByUUID error, %v", err)
		return nil
	}
	return users
}
