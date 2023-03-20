package entity

import (
	"productAccounting-v1/internal/domain/base"
)

type User struct {
	base.EntityWithGuidKey

	Password string `json:"password"`
	Email    string `json:"email" gorm:"uniqueIndex"`

	Role string `json:"role"`
}
