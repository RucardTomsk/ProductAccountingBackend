package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"productAccounting-v1/cmd/admin-api/api/controller"
	"productAccounting-v1/cmd/admin-api/config"
	"productAccounting-v1/internal/api/middleware"
	"productAccounting-v1/internal/common"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	config config.Config
}

func NewRouter(config config.Config) *Router {
	return &Router{
		config: config,
	}
}

func (r *Router) InitRoutes(
	logger *zap.Logger,
	container *controller.Container) *gin.Engine {

	gin.SetMode(r.config.Server.GinMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.SetTracingContext(*logger))
	router.Use(middleware.SetRequestLogging(*logger))
	router.Use(middleware.SetOperationName(r.config.Server, *logger))
	router.Use(cors.New(common.DefaultCorsConfig()))

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	v1 := baseRouter.Group("/v1")

	user := v1.Group("user")
	{
		user.POST("register", container.AuthController.AddUser)
		user.POST("login", container.AuthController.Login)
		user.POST("test", container.AuthController.MiddlewareCheckAdmin, container.AuthController.Test)
	}

	chapter := v1.Group("chapter")
	{
		chapter.POST("create", container.AuthController.MiddlewareCheckAdmin, container.ChapterController.CreateChapter)
		chapter.POST(":chapter-id/update", container.AuthController.MiddlewareCheckAdmin, container.ChapterController.UpdateChapter)
		chapter.POST(":chapter-id/subchapter/add", container.AuthController.MiddlewareCheckAdmin, container.ChapterController.CreateSubchapter)
		chapter.POST(":chapter-id/delete", container.AuthController.MiddlewareCheckAdmin, container.ChapterController.DeleteChapter)
		chapter.GET("get", container.AuthController.MiddlewareCheckAdmin, container.ChapterController.GetChapters)
		chapter.GET(":chapter-id/component/get", container.AuthController.MiddlewareCheckAdmin, container.ChapterController.GetComponents)
	}

	component := v1.Group("component")
	{
		component.POST("chapter/:chapter-id/create", container.AuthController.MiddlewareCheckAdmin, container.ComponentController.CreateComponent)
		component.POST(":component-id/add", container.AuthController.MiddlewareCheckAdmin, container.ComponentController.AddComponent)
		component.POST(":component-id/delete", container.AuthController.MiddlewareCheckAdmin, container.ComponentController.DeleteComponent)
		component.POST(":component-id/update", container.AuthController.MiddlewareCheckAdmin, container.ComponentController.UpdateComponent)
		component.POST(":component-id/use", container.AuthController.MiddlewareCheckAdmin, container.ComponentController.UseComponent)
		component.POST(":component-id/uploadImage", container.ComponentController.UploadImage)
	}

	product := v1.Group("product")
	{
		product.POST("create", container.AuthController.MiddlewareCheckAdmin, container.ProductController.CreateProduct)
		product.POST(":product-id/update", container.AuthController.MiddlewareCheckAdmin, container.ProductController.UpdateProduct)
		product.POST("assembly/:assembly-id/add", container.AuthController.MiddlewareCheckAdmin, container.ProductController.AddAssembly)
		product.POST("assembly/:assembly-id/component/component-id/add", container.AuthController.MiddlewareCheckAdmin, container.ProductController.AddComponentToAssembly)
		product.GET(":product-id/get", container.AuthController.MiddlewareCheckAdmin, container.ProductController.GetProduct)
		product.GET("assembly/:assembly-id/component/get", container.AuthController.MiddlewareCheckAdmin, container.ProductController.GetAssemblyComponent)
	}

	return router
}
