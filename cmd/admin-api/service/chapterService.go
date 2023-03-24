package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/storage/dao"
	"productAccounting-v1/internal/domain/base"
	"productAccounting-v1/internal/domain/entity"
)

type ChapterService struct {
	storage *dao.ChapterStorage
}

func NewChapterService(storage *dao.ChapterStorage) *ChapterService {
	return &ChapterService{
		storage: storage,
	}
}

func (s *ChapterService) CreateChapter(request *model.CreateChapterRequest) (*uuid.UUID, *base.ServiceError) {
	chapter := &entity.Chapter{
		Name: request.Name,
	}

	if err := s.storage.CreateChapter(chapter); err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return &chapter.ID, nil
}

func (s *ChapterService) UpdateChapter(chapterID *uuid.UUID, request *model.UpdateChapterRequest) *base.ServiceError {
	chapter, err := s.storage.RetrieveChapter(chapterID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	chapter.Name = request.Name

	if err := s.storage.UpdateChapter(chapter); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

func (s *ChapterService) AddSubchapter(chapterID *uuid.UUID, request *model.CreateChapterRequest) (*uuid.UUID, *base.ServiceError) {
	chapter, err := s.storage.RetrieveChapter(chapterID)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	subchapter := &entity.Chapter{
		Name:    request.Name,
		IsChild: true,
	}

	chapter.Subchapter = append(chapter.Subchapter, subchapter)

	if err := s.storage.UpdateChapter(chapter); err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	return &subchapter.ID, nil
}

func (s *ChapterService) DeleteChapter(id *uuid.UUID) *base.ServiceError {
	if err := s.storage.DeleteChapter(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	return nil
}
func (s *ChapterService) GetChapters() ([]model.ChapterObject, *base.ServiceError) {
	chapters, err := s.storage.GetChapters()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, base.NewNotFoundError(err)
		}
		return nil, base.NewPostgresReadError(err)
	}

	result := make([]model.ChapterObject, 0, len(chapters))

	for _, chapter := range chapters {
		subchapters := make([]model.ChapterObject, 0, len(chapter.Subchapter))
		for _, subchapter := range chapter.Subchapter {
			subchapters = append(subchapters, model.ChapterObject{
				ID:   subchapter.ID,
				Name: subchapter.Name,
			})
		}

		result = append(result, model.ChapterObject{
			ID:         chapter.ID,
			Name:       chapter.Name,
			Subchapter: subchapters,
		})
	}

	return result, nil
}

func (s *ChapterService) GetComponents(id *uuid.UUID) ([]model.ComponentObject, *base.ServiceError) {
	components, err := s.storage.GetComponents(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, base.NewNotFoundError(err)
		}
		return nil, base.NewPostgresReadError(err)
	}

	result := make([]model.ComponentObject, 0, len(components))

	for _, component := range components {
		result = append(result, model.ComponentObject{
			ID:     component.ID,
			Name:   component.Name,
			Price:  component.Price,
			Weight: component.Weight,
		})
	}

	return result, nil
}
