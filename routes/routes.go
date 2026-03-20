package routes

import (
	"template-go/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, envConfig *config.EnvConfig) {

	// versión API
	v1 := router.Group(envConfig.BasePath)

	// módulos
	RegisterUserRoutes(v1)
}
