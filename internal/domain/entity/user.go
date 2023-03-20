package entity

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/enum"
)

type User struct {
	ID uuid.UUID

	Password string
	Email    string

	Role enum.Roles
}
