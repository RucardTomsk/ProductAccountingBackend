package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/storage/dao"
	"productAccounting-v1/internal/domain/base"
	"productAccounting-v1/internal/domain/entity"
	"productAccounting-v1/internal/domain/enum"
)

type ComponentService struct {
	storage        *dao.ComponentStorage
	chapterStorage *dao.ChapterStorage
}

func NewComponentService(storage *dao.ComponentStorage,
	chapterStorage *dao.ChapterStorage) *ComponentService {
	return &ComponentService{
		storage:        storage,
		chapterStorage: chapterStorage,
	}
}

func (s *ComponentService) CreateComponent(chapterID *uuid.UUID, request *model.CreateComponentRequest) (*uuid.UUID, *base.ServiceError) {
	chapter, err := s.chapterStorage.RetrieveChapter(chapterID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, base.NewNotFoundError(err)
		}

		return nil, base.NewPostgresReadError(err)
	}

	component := &entity.Component{
		Name:      request.Name,
		Chapter:   chapter,
		ChapterID: *chapterID,
	}

	if err := s.storage.CreateComponent(component); err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return &component.ID, nil
}

func (s *ComponentService) AddComponent(componentID *uuid.UUID, request *model.UpdateComponent) *base.ServiceError {
	component, err := s.storage.RetrieveComponent(componentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	if enum.ParseTypeWeight(request.TypeWeight) == enum.KG {
		priseToG := request.Price / (request.Weight * 1000)
		component.Weight = component.Weight + request.Weight*1000
		component.Price = component.Weight * priseToG
	} else {
		priseToG := request.Price / request.Weight
		component.Weight = component.Weight + request.Weight
		component.Price = component.Weight * priseToG
	}

	if err := s.storage.UpdateComponent(component); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

func (s *ComponentService) DeleteComponent(id *uuid.UUID) *base.ServiceError {
	if err := s.storage.DeleteComponent(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresWriteError(err)
	}

	return nil
}