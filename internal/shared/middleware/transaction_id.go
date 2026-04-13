package middleware

import (
	"app/internal/shared/config/logger"
	"app/internal/shared/config/resources"
	"app/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

func TransactionID() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.GetHeader(resources.TransactionIDKey)

		if transactionID == "" {
			transactionID = utils.GenerateTransactionID()
		}

		ctx := logger.WithTransactionID(c.Request.Context(), transactionID)
		c.Request = c.Request.WithContext(ctx)

		c.Set(resources.TransactionIDKey, transactionID)

		c.Next()
	}
}
