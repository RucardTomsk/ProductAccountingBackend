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

type ChapterController struct {
	logger  *zap.Logger
	service *service.ChapterService
}

func NewChapterController(
	logger *zap.Logger,
	service *service.ChapterService) *ChapterController {
	return &ChapterController{
		logger:  logger,
		service: service,
	}
}

func (a *ChapterController) CreateChapter(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	var payload model.CreateChapterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	id, serviceErr := a.service.CreateChapter(&payload)
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

func (a *ChapterController) UpdateChapter(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	chapterID, err := uuid.Parse(c.Params.ByName("chapter-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.UpdateChapterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.UpdateChapter(&chapterID, &payload); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOK{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

func (a *ChapterController) CreateSubchapter(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	chapterID, err := uuid.Parse(c.Params.ByName("chapter-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.CreateChapterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	id, serviceErr := a.service.AddSubchapter(&chapterID, &payload)
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

func (a *ChapterController) DeleteChapter(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	chapterID, err := uuid.Parse(c.Params.ByName("chapter-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.DeleteChapter(&chapterID); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOK{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

func (a *ChapterController) GetChapters(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	chapters, serviceErr := a.service.GetChapters()
	if serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.GetChaptersResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Chapters: chapters,
	})
}

func (a *ChapterController) GetComponents(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	chapterID, err := uuid.Parse(c.Params.ByName("chapter-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	components, serviceErr := a.service.GetComponents(&chapterID)
	if serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.GetComponentsResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Components: components,
	})
}
