package routes

import (
	app_routes "app/internal/app/interfaces/routes"
	config "app/internal/shared/config/env"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, envConfig *config.EnvConfig) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group(envConfig.ServicePrefix)

	app_routes.RegisterUserRoutes(apiGroup, envConfig)
}
