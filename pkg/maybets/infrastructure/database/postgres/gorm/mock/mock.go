package mock

import (
	"context"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	"github.com/google/uuid"
)

// GormMock mocks caching implementations
type GormMock struct {
	MockGetTotalBetsFn      func(ctx context.Context, userID string) (int64, error)
	MockGetTotalWinningsFn  func(ctx context.Context, userID string) (float64, error)
	MockGetTopUsersFn       func(ctx context.Context, limit int) ([]gorm.User, error)
	MockGetAnomalousUsersFn func(ctx context.Context) ([]gorm.User, error)
	MockStoreBetDataFn      func(ctx context.Context, bets []gorm.Bet) error
}

// NewGormMock initializes our client mocks
func NewGormMock() *GormMock {
	return &GormMock{
		MockGetTotalBetsFn: func(_ context.Context, _ string) (int64, error) {
			return 5, nil
		},
		MockGetTotalWinningsFn: func(_ context.Context, _ string) (float64, error) {
			return 100.00, nil
		},
		MockGetTopUsersFn: func(_ context.Context, _ int) ([]gorm.User, error) {
			return []gorm.User{
				{
					UserID:    uuid.NewString(),
					TotalBets: 4,
				},
			}, nil
		},
		MockGetAnomalousUsersFn: func(_ context.Context) ([]gorm.User, error) {
			return []gorm.User{
				{
					UserID:    uuid.NewString(),
					TotalBets: 4,
				},
			}, nil
		},
		MockStoreBetDataFn: func(_ context.Context, _ []gorm.Bet) error {
			return nil
		},
	}
}

// GetTotalBets mocks retrieval of total bets
func (g *GormMock) GetTotalBets(ctx context.Context, userID string) (int64, error) {
	return g.MockGetTotalBetsFn(ctx, userID)
}

// GetTotalWinnings mocks retrieval of a user's total winnings
func (g *GormMock) GetTotalWinnings(ctx context.Context, userID string) (float64, error) {
	return g.MockGetTotalWinningsFn(ctx, userID)
}

// GetTopUsers mocks retrieval of the users with most bets
func (g *GormMock) GetTopUsers(ctx context.Context, limit int) ([]gorm.User, error) {
	return g.MockGetTopUsersFn(ctx, limit)
}

// GetAnomalousUsers mocks retrieval of the users with high betting activity
func (g *GormMock) GetAnomalousUsers(ctx context.Context) ([]gorm.User, error) {
	return g.MockGetAnomalousUsersFn(ctx)
}

// StoreBetData mocks storing user bet data
func (g *GormMock) StoreBetData(ctx context.Context, bets []gorm.Bet) error {
	return g.MockStoreBetDataFn(ctx, bets)
}
