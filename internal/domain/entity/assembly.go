package entity

import "productAccounting-v1/internal/domain/base"

type Assembly struct {
	base.EntityWithGuidKey

	Name       string
	Weight     float64     `json:"weight"`
	Components []Component `json:"components,omitempty" gorm:"many2many:assembly_components"`
}
