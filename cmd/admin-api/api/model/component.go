package model

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type (
	CreateComponentRequest struct {
		Name string `json:"name"`
	}

	UpdateComponent struct {
		Name       string  `json:"name"`
		Weight     float64 `json:"weight"`
		Price      float64 `json:"price"`
		TypeWeight string  `json:"typeWeight"`
	}

	ComponentObject struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Weight float64   `json:"weight"`
		Price  float64   `json:"price"`
	}

	UseComponentRequest struct {
		Weight float64 `json:"weight"`
	}

	GetComponentsResponse struct {
		base.ResponseOK
		Components []ComponentObject `json:"components"`
	}
)
