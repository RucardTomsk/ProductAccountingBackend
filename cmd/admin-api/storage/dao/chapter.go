package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"productAccounting-v1/internal/domain/entity"
)

type ChapterStorage struct {
	db *gorm.DB
}

func NewChapterStorage(db *gorm.DB) *ChapterStorage {
	return &ChapterStorage{
		db: db,
	}
}

func (s *ChapterStorage) CreateChapter(chapter *entity.Chapter) error {
	return s.db.Create(&chapter).Error
}

func (s *ChapterStorage) UpdateChapter(chapter *entity.Chapter) error {
	tx := s.db.Updates(chapter)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *ChapterStorage) RetrieveChapter(id *uuid.UUID) (*entity.Chapter, error) {
	var chapter entity.Chapter
	err := s.db.First(&chapter, id).Error
	return &chapter, err
}

func (s *ChapterStorage) GetChapters() ([]entity.Chapter, error) {
	var chapters []entity.Chapter
	err := s.db.Model(&entity.Chapter{}).
		Preload("Subchapter").
		Where("is_child = ?", false).
		Find(&chapters).Error

	return chapters, err
}

func (s *ChapterStorage) GetComponents(id *uuid.UUID) ([]entity.Component, error) {
	var chapter entity.Chapter
	err := s.db.Model(&entity.Chapter{}).
		Preload("Components").
		First(&chapter, id).Error

	return chapter.Components, err
}

func (s *ChapterStorage) DeleteChapter(id *uuid.UUID) error {
	tx := s.db.Delete(&entity.Chapter{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
