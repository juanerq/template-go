package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// versión API
	v1 := router.Group("/api/v1")

	// módulos
	RegisterUserRoutes(v1)
}
