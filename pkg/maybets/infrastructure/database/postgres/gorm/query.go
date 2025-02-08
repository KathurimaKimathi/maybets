package gorm

import (
	"context"
	"fmt"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/enums"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var tracer = otel.Tracer("github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm")

// GetTotalBets fetches the total number of bets placed by a user.
func (db DBInstance) GetTotalBets(ctx context.Context, userID string) (int64, error) {
	_, span := tracer.Start(ctx, "GetTotalBets")
	defer span.End()

	var totalBets int64
	err := db.DB.Model(&Bet{}).
		Where("user_id = ?", userID).
		Count(&totalBets).Error

	if err != nil {
		span.SetStatus(codes.Error, "Failed to fetch total bets")
		span.RecordError(err)

		return 0, fmt.Errorf("failed to get total bets: %w", err)
	}

	return totalBets, nil
}

// GetTotalWinnings calculates the total winnings of a user.
func (db DBInstance) GetTotalWinnings(ctx context.Context, userID string) (float64, error) {
	_, span := tracer.Start(ctx, "GetTotalWinnings")
	defer span.End()

	var totalWinnings float64
	err := db.DB.Model(&Bet{}).
		Where("user_id = ? AND outcome = ?", userID, enums.Win).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalWinnings).Error

	if err != nil {
		span.SetStatus(codes.Error, "Failed to calculate total winnings")
		span.RecordError(err)

		return 0, fmt.Errorf("failed to get total winnings: %w", err)
	}

	return totalWinnings, nil
}

// GetTopUsers fetches the top users with the highest betting volume.
func (db DBInstance) GetTopUsers(ctx context.Context, limit int) ([]User, error) {
	_, span := tracer.Start(ctx, "GetTopUsers")
	defer span.End()

	var topUsers []User
	err := db.DB.Model(&Bet{}).
		Select("user_id, COUNT(*) as total_bets").
		Group("user_id").
		Order("total_bets DESC").
		Limit(limit).
		Scan(&topUsers).Error

	if err != nil {
		span.SetStatus(codes.Error, "Failed to fetch top users")
		span.RecordError(err)

		return nil, fmt.Errorf("failed to get top users: %w", err)
	}

	return topUsers, nil
}

// GetAnomalousUsers fetches users with significantly higher betting activity than the average.
func (db DBInstance) GetAnomalousUsers(ctx context.Context) ([]User, error) {
	_, span := tracer.Start(ctx, "GetAnomalousUsers")
	defer span.End()

	// get the total number of bets and the total number of distinct users
	var totalBets, totalUsers int64

	err := db.DB.Model(&Bet{}).
		Select("COUNT(*)").
		Scan(&totalBets).Error
	if err != nil {
		span.SetStatus(codes.Error, "Failed to count total bets")
		span.RecordError(err)

		return nil, fmt.Errorf("failed to count total bets: %w", err)
	}

	err = db.DB.Model(&Bet{}).
		Distinct("user_id").
		Count(&totalUsers).Error
	if err != nil {
		span.SetStatus(codes.Error, "Failed to count distinct users")
		span.RecordError(err)

		return nil, fmt.Errorf("failed to count distinct users: %w", err)
	}

	// average number of bets per user
	avgBets := float64(totalBets) / float64(totalUsers)

	// users who have placed bets significantly above the average eg 2 times the average
	var anomalousUsers []User

	err = db.DB.Model(&Bet{}).
		Select("user_id, COUNT(*) as total_bets").
		Group("user_id").
		Having("COUNT(*) > ?", avgBets*2.5).
		Order("total_bets DESC").
		Scan(&anomalousUsers).Error
	if err != nil {
		span.SetStatus(codes.Error, "Failed to fetch anomalous users")
		span.RecordError(err)

		return nil, fmt.Errorf("failed to get anomalous users: %w", err)
	}

	return anomalousUsers, nil
}
