package dao

import (
	"gorm.io/gorm"
	"productAccounting-v1/internal/domain/entity"
)

type AuthStorage struct {
	db *gorm.DB
}

func NewAuthStorage(db *gorm.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) CreateUser(user *entity.User) error {
	return s.db.Create(&user).Error
}

func (s *AuthStorage) GetUser(email string, password string) (*entity.User, error) {
	var user entity.User
	tx := s.db.Model(entity.User{}).
		Where("password = ? AND email = ?", password, email).
		First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
