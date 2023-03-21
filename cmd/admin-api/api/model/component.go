package model

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type (
	CreateComponentRequest struct {
		Name string
	}

	UpdateComponent struct {
		Name       string
		Weight     float64
		Price      float64
		TypeWeight string
	}

	ComponentObject struct {
		ID     uuid.UUID
		Name   string
		Weight float64
		Price  float64
	}

	UseComponentRequest struct {
		Weight float64
	}

	GetComponentsResponse struct {
		base.ResponseOK
		Components []ComponentObject
	}
)
