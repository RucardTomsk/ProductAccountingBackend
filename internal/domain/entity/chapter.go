package entity

import (
	"productAccounting-v1/internal/domain/base"
)

type Chapter struct {
	base.EntityWithGuidKey
	Name string `json:"name"`

	Chapters   []Chapter   `gorm:"many2many:chapter_chapter;"`
	Components []Component `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
