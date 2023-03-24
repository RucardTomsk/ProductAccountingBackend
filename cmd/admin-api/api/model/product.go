package model

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type (
	CreateProductRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	UpdateProductRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	CreateAssemblyRequest struct {
		Name   string  `json:"name"`
		Weight float64 `json:"weight"`
	}

	AssemblyObject struct {
		ID         uuid.UUID         `json:"id"`
		Name       string            `json:"name"`
		Weight     float64           `json:"weight"`
		Components []ComponentObject `json:"components"`
	}

	ProductObject struct {
		ID          uuid.UUID        `json:"id"`
		Name        string           `json:"name"`
		Description string           `json:"description"`
		Assembles   []AssemblyObject `json:"assembles"`
	}

	GetAssemblesResponse struct {
		base.ResponseOK
		Assembles []AssemblyObject `json:"assembles"`
	}

	GetProductResponse struct {
		base.ResponseOK
		Products []ProductObject `json:"products"`
	}
)
