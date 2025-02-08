package postgres

import (
	"context"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
)

// Cache interface holds methods for interacting with the caching service
type Cache interface {
	Get(ctx context.Context, key string, valueType interface{}) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

// Query holds the method signatures used to query the database
type Query interface {
	GetTotalBets(ctx context.Context, userID string) (int64, error)
	GetTotalWinnings(ctx context.Context, userID string) (float64, error)
	GetTopUsers(ctx context.Context, limit int) ([]gorm.User, error)
	GetAnomalousUsers(ctx context.Context) ([]gorm.User, error)
}

// Create contains the method signatures used to create a new record in the database
type Create interface {
	StoreBetData(ctx context.Context, bet []gorm.Bet) error
}

// MaybetsDB struct implements the service's business specific calls to the database
type MaybetsDB struct {
	cache  Cache
	query  Query
	create Create
}

// NewMaybetsDB initializes a new instance of the MaybetsDB struct
func NewMaybetsDB(c Cache, q Query, cr Create) *MaybetsDB {
	return &MaybetsDB{
		cache:  c,
		query:  q,
		create: cr,
	}
}
