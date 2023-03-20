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
	"productAccounting-v1/internal/domain/enum"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userGuid            = "userGuid"
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

func (a *AuthController) MiddlewareCheckAdmin(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, serviceErr := a.service.ParseToken(headerParts[1])
	if serviceErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, serviceErr)
		return
	}

	if enum.ParseRoles(claims.UserRole) != enum.ADMIN {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	c.Next()
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

func (a *AuthController) Login(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	var payload model.AuthRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	token, serviceErr := a.service.Login(&payload)

	if serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithJWT{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
		JWT:        token.Value,
	})
}
