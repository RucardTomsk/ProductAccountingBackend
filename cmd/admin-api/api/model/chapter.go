package model

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type (
	CreateChapterRequest struct {
		Name string `json:"name"`
	}

	UpdateChapterRequest struct {
		Name string `json:"name"`
	}

	ChapterObject struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`

		Subchapter []ChapterObject   `json:"subchapter"`
		Components []ComponentObject `json:"components"`
	}

	GetChaptersResponse struct {
		base.ResponseOK
		Chapters []ChapterObject `json:"chapters"`
	}
)
