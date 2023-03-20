package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"productAccounting-v1/cmd/admin-api/api/controller"
	"productAccounting-v1/cmd/admin-api/api/router"
	"productAccounting-v1/cmd/admin-api/config"
	"productAccounting-v1/cmd/admin-api/docs"
	"productAccounting-v1/cmd/admin-api/service"
	"productAccounting-v1/cmd/admin-api/storage/dao"
	"productAccounting-v1/cmd/admin-api/storage/driver"
	"productAccounting-v1/internal/common"
	"productAccounting-v1/internal/server"
	"productAccounting-v1/internal/telemetry/log"

	"syscall"
)

func main() {
	logger := log.NewLogger()

	appCli := common.InitAppCli()
	if err := appCli.Run(os.Args); err != nil {
		logger.Fatal(err.Error())
	}

	// read config
	var cfg config.Config
	if err := viper.MergeInConfig(); err != nil {
		logger.Fatal(fmt.Sprintf("error reading config file: %v", err))
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to decode into struct: %v", err))
	}

	// configure swagger
	swaggerConfig := common.NewSwaggerConfig("admin api", "TBD", "unreleased")

	docs.SwaggerInfo.Title = swaggerConfig.Title
	docs.SwaggerInfo.Description = swaggerConfig.Description
	docs.SwaggerInfo.Version = swaggerConfig.Version
	docs.SwaggerInfo.Host = swaggerConfig.Host
	docs.SwaggerInfo.BasePath = swaggerConfig.BasePath
	docs.SwaggerInfo.Schemes = swaggerConfig.Schemes

	neoDriver, err := driver.NewNeo4jDriver(&cfg.Neo4j)

	if err != nil {
		logger.Fatal(fmt.Sprintf("can't connect to database: %v", err))
	}

	// init storage
	authStorage := dao.NewAuthStorage(neoDriver)

	// init services
	authService := service.NewAuthService(authStorage)

	// init controllers
	controllers := controller.NewControllerContainer(
		logger,
		authService)

	handler := router.NewRouter(cfg)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Server.Host, cfg.Server.Port, handler.InitRoutes(
			logger,
			authService,
			controllers)); err != nil {
			logger.Error(fmt.Sprintf("error accured while running http server: %s", err.Error()))
		}
	}()

	logger.Info(fmt.Sprintf("listening on %s:%s", cfg.Server.Host, cfg.Server.Port))

	// handle signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("shutting down gracefully...")
	defer func() { logger.Info("shutdown complete") }()

	// perform shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error(fmt.Sprintf("error occured on server shutting down: %s", err.Error()))
	}
}
