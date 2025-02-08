package infrastructure

import (
	"context"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
)

// Database holds the methods of interacting with the database
type Database interface {
	GetTotalBets(ctx context.Context, userID string) (int64, error)
	GetTotalWinnings(ctx context.Context, userID string) (float64, error)
	GetTopUsers(ctx context.Context, limit int) ([]domain.User, error)
	GetAnomalousUsers(ctx context.Context) ([]domain.User, error)
	StoreBetData(ctx context.Context, bets []*domain.Bet) error
}

// Cache interface holds methods for interacting with the caching service
type Cache interface {
	Get(ctx context.Context, key string, valueType interface{}) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

// Infrastructure implements the infrastructure interface(s)
type Infrastructure struct {
	Cache    Cache
	Database Database
}

// NewInfrastructureInteractor initializes a new Infrastructure
func NewInfrastructureInteractor(
	cache Cache,
	database Database,
) *Infrastructure {
	return &Infrastructure{
		Cache:    cache,
		Database: database,
	}
}
