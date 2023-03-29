package service

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"math"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/storage/dao"
	"productAccounting-v1/internal/domain/base"
	"productAccounting-v1/internal/domain/entity"
	"productAccounting-v1/internal/domain/enum"
	"productAccounting-v1/internal/s3"
)

type ComponentService struct {
	storage        *dao.ComponentStorage
	chapterStorage *dao.ChapterStorage
	minioService   *s3.MinioService
}

func NewComponentService(storage *dao.ComponentStorage,
	chapterStorage *dao.ChapterStorage,
	minioService *s3.MinioService) *ComponentService {
	return &ComponentService{
		storage:        storage,
		chapterStorage: chapterStorage,
		minioService:   minioService,
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
		Name:        request.Name,
		Description: request.Description,
		Chapter:     chapter,
		ChapterID:   *chapterID,
	}

	if err := s.storage.CreateComponent(component); err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return &component.ID, nil
}

func (s *ComponentService) UpdateComponent(componentID *uuid.UUID, request *model.UpdateComponent) *base.ServiceError {
	component, err := s.storage.RetrieveComponent(componentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	component.Name = request.Name
	component.Description = request.Description

	if err := s.storage.UpdateComponent(component); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
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
		component.Price = math.Ceil((request.Price+(request.Price*0.15))/(request.Weight*1000)*100) / 100
		component.Weight = component.Weight + request.Weight*1000
	} else {
		component.Price = math.Ceil((request.Price+(request.Price*0.15))/request.Weight*100) / 100
		component.Weight = component.Weight + request.Weight
	}

	if err := s.storage.UpdateComponent(component); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

func (s *ComponentService) UseComponent(componentID *uuid.UUID, request *model.UseComponentRequest) *base.ServiceError {
	component, err := s.storage.RetrieveComponent(componentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	component.Weight = component.Weight - request.Weight
	component.Price = component.Weight * component.Price

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

func (s *ComponentService) UploadImage(ctx context.Context, file io.Reader, size int64, contentType, nameFile string) *base.ServiceError {
	if err := s.minioService.Upload(ctx, s3.UploadInput{
		File:        file,
		Size:        size,
		ContentType: contentType,
		Name:        nameFile,
	}); err != nil {
		return err
	}

	return nil
}
