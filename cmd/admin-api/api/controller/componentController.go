package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/service"
	"productAccounting-v1/internal/api"
	"productAccounting-v1/internal/api/middleware"
	"productAccounting-v1/internal/domain/base"
)

type ComponentController struct {
	logger  *zap.Logger
	service *service.ComponentService
}

func NewComponentController(
	logger *zap.Logger,
	service *service.ComponentService) *ComponentController {
	return &ComponentController{
		service: service,
		logger:  logger,
	}
}

func (a *ComponentController) CreateComponent(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	chapterID, err := uuid.Parse(c.Params.ByName("chapter-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.CreateComponentRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	id, serviceErr := a.service.CreateComponent(&chapterID, &payload)
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

func (a *ComponentController) AddComponent(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	componentID, err := uuid.Parse(c.Params.ByName("component-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.UpdateComponent
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.AddComponent(&componentID, &payload); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

func (a *ComponentController) DeleteComponent(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	componentID, err := uuid.Parse(c.Params.ByName("component-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.DeleteComponent(&componentID); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}
