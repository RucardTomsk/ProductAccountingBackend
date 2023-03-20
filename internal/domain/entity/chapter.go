package entity

import (
	"productAccounting-v1/internal/domain/base"
)

type Chapter struct {
	base.EntityWithGuidKey
	Name string `json:"name"`

	IsChild    bool        `json:"isChild" gorm:"default:false;"`
	Subchapter []*Chapter  `gorm:"many2many:chapter_subchapter"`
	Components []Component `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
