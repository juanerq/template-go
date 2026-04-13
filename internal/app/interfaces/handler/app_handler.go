package app_handler

import (
	"app/internal/shared/config/logger"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Obtener estado del servicio
// @Description Retorna el estado del microservicio
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /MS/COM/Project/Service/V1/api [get]
func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "GetUsers",
		"transactionId": logger.GetTransactionID(c.Request.Context()),
	})
}
