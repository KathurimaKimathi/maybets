package gorm

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/codes"
)

// StoreBetData is used to store bet records in the database
func (db DBInstance) StoreBetData(ctx context.Context, bet []Bet) error {
	_, span := tracer.Start(ctx, "StoreBetData")
	defer span.End()

	err := db.DB.Create(&bet).Error
	if err != nil {
		span.SetStatus(codes.Error, "Failed to store bet data")
		span.RecordError(err)

		return fmt.Errorf("failed to store bet data: %w", err)
	}

	return nil
}
