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

type ProductController struct {
	logger  *zap.Logger
	service *service.ProductService
}

func NewProductController(
	logger *zap.Logger,
	service *service.ProductService) *ProductController {
	return &ProductController{
		service: service,
		logger:  logger,
	}
}

func (a *ProductController) CreateProduct(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	var payload model.CreateProductRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	id, serviceErr := a.service.CreateProduct(&payload)
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

func (a *ProductController) UpdateProduct(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	productID, err := uuid.Parse(c.Params.ByName("product-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.UpdateProductRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.UpdateProduct(&productID, &payload); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

func (a *ProductController) AddAssembly(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	productID, err := uuid.Parse(c.Params.ByName("product-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.CreateAssemblyRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Warn("error parsing json:" + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.AddAssembly(&productID, &payload); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

func (a *ProductController) AddComponentToAssembly(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	assemblyID, err := uuid.Parse(c.Params.ByName("assembly-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	componentID, err := uuid.Parse(c.Params.ByName("component-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.service.AddComponentToAssembly(&assemblyID, &componentID); serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

func (a *ProductController) GetProduct(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	products, serviceErr := a.service.GetProduct()
	if serviceErr != nil {
		log.Warn("error occurred: " + serviceErr.Error())
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.GetProductResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Products: products,
	})
}

func (a *ProductController) GetAssemblyComponent(c *gin.Context) {
	log := api.EnrichLogger(a.logger, c)

	assemblyID, err := uuid.Parse(c.Params.ByName("assembly-id"))
	if err != nil {
		log.Warn("error parsing uuid: " + err.Error())
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	components, serviceErr := a.service.GetAssemblyComponent(&assemblyID)
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
