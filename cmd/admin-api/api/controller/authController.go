package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/service"
	"productAccounting-v1/internal/api"
	"productAccounting-v1/internal/api/middleware"
	"productAccounting-v1/internal/domain/base"
)

type AuthController struct {
	logger  *zap.Logger
	service *service.AuthService
}

func NewAuthController(
	logger *zap.Logger,
	service *service.AuthService) *AuthController {
	return &AuthController{
		logger:  logger,
		service: service,
	}
}

func (a *AuthController) AddUser(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	var payload model.RegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	id, serviceErr := a.service.Register(&payload)
	if serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
		ID:         *id,
	})
}
