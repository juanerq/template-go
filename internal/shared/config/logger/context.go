package logger

import (
	"app/internal/shared/config/resources"
	"context"
)

func WithTransactionID(ctx context.Context, transactionID string) context.Context {
	return context.WithValue(ctx, resources.TransactionIDKey, transactionID)
}

func GetTransactionID(ctx context.Context) string {
	if txID, ok := ctx.Value(resources.TransactionIDKey).(string); ok && txID != "" {
		return txID
	}
	return "-"
}
