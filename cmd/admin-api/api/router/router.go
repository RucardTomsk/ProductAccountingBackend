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
	}

	return router
}
