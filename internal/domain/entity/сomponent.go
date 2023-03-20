package entity

import (
	"productAccounting-v1/internal/domain/base"
)

type Component struct {
	base.EntityWithGuidKey
	Name         string
	Weight       int
	Price        float64
	Availability bool
}
