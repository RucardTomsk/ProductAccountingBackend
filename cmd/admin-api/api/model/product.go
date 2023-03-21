package model

import (
	"github.com/google/uuid"
	"productAccounting-v1/internal/domain/base"
)

type (
	CreateProductRequest struct {
		Name        string
		Description string
	}

	UpdateProductRequest struct {
		Name        string
		Description string
	}

	CreateAssemblyRequest struct {
		Name   string
		Weight float64
	}

	AssemblyObject struct {
		ID         uuid.UUID
		Name       string
		Weight     float64
		Components []ComponentObject
	}

	ProductObject struct {
		ID          uuid.UUID
		Name        string
		Description string
		Assembles   []AssemblyObject
	}

	GetAssemblesResponse struct {
		base.ResponseOK
		Assembles []AssemblyObject
	}

	GetProductResponse struct {
		base.ResponseOK
		Products []ProductObject
	}
)
