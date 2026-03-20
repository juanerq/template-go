package routes

import (
	"template-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/users")

	users.GET("", handlers.GetUsers)
}
