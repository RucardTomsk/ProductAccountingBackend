package entity

import (
	"productAccounting-v1/internal/domain/base"
)

type Product struct {
	base.EntityWithGuidKey
	Name string `json:"name"`
}
