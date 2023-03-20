package model

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type (
	CreateChapterRequest struct {
		Name string
	}

	UpdateChapterRequest struct {
		Name string
	}

	ChapterObject struct {
		ID   uuid.UUID
		Name string

		Subchapter []ChapterObject
		Components []ComponentObject
	}

	GetChaptersResponse struct {
		base.ResponseOK
		Chapters []ChapterObject
	}
)
