package controller

import (
	"go.uber.org/zap"
	"productAccounting-v1/cmd/admin-api/service"
)

type Container struct {
	AuthController *AuthController
}

func NewControllerContainer(
	logger *zap.Logger,
	authService *service.AuthService) *Container {
	return &Container{
		AuthController: NewAuthController(logger, authService),
	}
}
