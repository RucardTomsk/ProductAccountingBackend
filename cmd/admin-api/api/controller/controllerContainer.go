package controller

import (
	"go.uber.org/zap"
	"productAccounting-v1/cmd/admin-api/service"
)

type Container struct {
	AuthController      *AuthController
	ChapterController   *ChapterController
	ComponentController *ComponentController
}

func NewControllerContainer(
	logger *zap.Logger,
	authService *service.AuthService,
	chapterService *service.ChapterService,
	componentService *service.ComponentService) *Container {
	return &Container{
		AuthController:      NewAuthController(logger, authService),
		ChapterController:   NewChapterController(logger, chapterService),
		ComponentController: NewComponentController(logger, componentService),
	}
}
