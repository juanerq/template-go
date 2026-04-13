package utils

import (
	"github.com/google/uuid"
)

// GenerateTransactionID genera un UUID v4 único
func GenerateTransactionID() string {
	return uuid.New().String()
}
