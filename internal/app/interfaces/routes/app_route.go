package app_routes

import (
	app_handler "app/internal/app/interfaces/handler"
	config "app/internal/shared/config/env"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, envConfig *config.EnvConfig) {

	rg.GET(envConfig.ControllerPath, app_handler.GetUsers)
}
