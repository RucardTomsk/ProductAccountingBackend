package controller

import (
	"go.uber.org/zap"
	"productAccounting-v1/cmd/admin-api/service"
)

type Container struct {
	AuthController      *AuthController
	ChapterController   *ChapterController
	ComponentController *ComponentController
	ProductController   *ProductController
}

func NewControllerContainer(
	logger *zap.Logger,
	authService *service.AuthService,
	chapterService *service.ChapterService,
	componentService *service.ComponentService,
	productService *service.ProductService) *Container {
	return &Container{
		AuthController:      NewAuthController(logger, authService),
		ChapterController:   NewChapterController(logger, chapterService),
		ComponentController: NewComponentController(logger, componentService),
		ProductController:   NewProductController(logger, productService),
	}
}
