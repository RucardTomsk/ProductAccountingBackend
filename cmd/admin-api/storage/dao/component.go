package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"productAccounting-v1/internal/domain/entity"
)

type ComponentStorage struct {
	db *gorm.DB
}

func NewComponentStorage(db *gorm.DB) *ComponentStorage {
	return &ComponentStorage{
		db: db,
	}
}

func (s *ComponentStorage) CreateComponent(component *entity.Component) error {
	return s.db.Create(&component).Error
}

func (s *ComponentStorage) UpdateComponent(component *entity.Component) error {
	tx := s.db.Updates(component)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *ComponentStorage) RetrieveComponent(id *uuid.UUID) (*entity.Component, error) {
	var component entity.Component
	err := s.db.First(&component, id).Error
	return &component, err
}

func (s *ComponentStorage) DeleteComponent(id *uuid.UUID) error {
	tx := s.db.Delete(&entity.Component{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
