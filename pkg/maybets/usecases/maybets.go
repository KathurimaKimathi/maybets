package usecases

import (
	"context"
	"log"
	"sync"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/KathurimaKimathi/maybets/pkg/maybets/usecases/")

// GetUserTotalBets fetches the total number of bets placed by a user.
func (u *UsecaseMayBets) GetUserTotalBets(ctx context.Context, userID string) (*domain.User, error) {
	_, span := tracer.Start(ctx, "GetUserTotalBets")
	defer span.End()

	totalBets, err := u.Infrastructure.Database.GetTotalBets(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        userID,
		TotalBets: totalBets,
	}, nil
}

// GetUserTotalWinnings calculates the total winnings of a user.
func (u *UsecaseMayBets) GetUserTotalWinnings(ctx context.Context, userID string) (*domain.User, error) {
	_, span := tracer.Start(ctx, "GetUserTotalWinnings")
	defer span.End()

	totalWinnings, err := u.Infrastructure.Database.GetTotalWinnings(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:            userID,
		TotalWinnings: totalWinnings,
	}, nil
}

// GetTopFiveUsers fetches the top 5 users with the highest betting volume.
func (u *UsecaseMayBets) GetTopFiveUsers(ctx context.Context) ([]domain.User, error) {
	_, span := tracer.Start(ctx, "GetTopFiveUsers")
	defer span.End()

	users, err := u.Infrastructure.Database.GetTopUsers(ctx, 5)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetAllAnomalousUsers fetches users with significantly higher betting activity than the average.
func (u *UsecaseMayBets) GetAllAnomalousUsers(ctx context.Context) ([]domain.User, error) {
	_, span := tracer.Start(ctx, "GetAllAnomalousUsers")
	defer span.End()

	users, err := u.Infrastructure.Database.GetAnomalousUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// ProcessBets processes bets concurrently and in batches of 1000
func (u *UsecaseMayBets) ProcessBets(ctx context.Context, bets []*domain.Bet) error {
	var wg sync.WaitGroup

	batchSize := 1000

	for i := 0; i < len(bets); i += batchSize {
		wg.Add(1)

		go func(batch []*domain.Bet) {
			defer wg.Done()

			if err := u.Infrastructure.Database.StoreBetData(ctx, batch); err != nil {
				log.Printf("Error processing batch: %v", err)
			}
		}(bets[i:min(i+batchSize, len(bets))])
	}

	wg.Wait()

	return nil
}
