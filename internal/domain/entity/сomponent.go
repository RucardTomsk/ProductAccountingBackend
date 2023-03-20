package entity

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type Component struct {
	base.EntityWithGuidKey
	Name   string  `json:"name"`
	Weight float64 `json:"weight" gorm:"default:0;"`
	Price  float64 `json:"price" gorm:"default:0;"`

	ChapterID uuid.UUID `json:"chapterID"`
	Chapter   *Chapter  `json:"chapter,omitempty"`
}
