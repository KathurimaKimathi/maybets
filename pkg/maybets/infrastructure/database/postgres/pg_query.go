package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/")

// GetTotalBets fetches the total number of bets placed by a user.
func (db MaybetsDB) GetTotalBets(ctx context.Context, userID string) (int64, error) {
	_, span := tracer.Start(ctx, "GetTotalBets")
	defer span.End()

	cacheKey := fmt.Sprintf("total-bets-%s", userID)

	cachedTotal, err := db.cache.Get(ctx, cacheKey, new(*int64))
	if err == nil {
		totalInt, ok := cachedTotal.(*int64)
		if !ok {
			return 0, fmt.Errorf("cannot cast interface to int64 pointer type")
		}

		return *totalInt, nil
	}

	fetchedTotal, err := db.query.GetTotalBets(ctx, userID)
	if err != nil {
		return 0, err
	}

	err = db.cache.Set(ctx, cacheKey, &fetchedTotal, time.Minute)
	if err != nil {
		log.Println(err.Error())
	}

	return fetchedTotal, nil
}

// GetTotalWinnings calculates the total winnings of a user.
func (db MaybetsDB) GetTotalWinnings(ctx context.Context, userID string) (float64, error) {
	_, span := tracer.Start(ctx, "GetTotalWinnings")
	defer span.End()

	cacheKey := fmt.Sprintf("total-winnings-%s", userID)

	cachedTotal, err := db.cache.Get(ctx, cacheKey, new(*float64))
	if err == nil {
		totalFloat, ok := cachedTotal.(*float64)
		if !ok {
			return 0, fmt.Errorf("cannot cast interface to float64 pointer type")
		}

		return *totalFloat, nil
	}

	fetchedTotal, err := db.query.GetTotalWinnings(ctx, userID)
	if err != nil {
		return 0, err
	}

	err = db.cache.Set(ctx, cacheKey, &fetchedTotal, time.Minute)
	if err != nil {
		log.Println(err.Error())
	}

	return fetchedTotal, nil
}

// GetTopUsers fetches the top users with the highest betting volume.
func (db MaybetsDB) GetTopUsers(ctx context.Context, limit int) ([]domain.User, error) {
	_, span := tracer.Start(ctx, "GetTopUsers")
	defer span.End()

	cacheKey := fmt.Sprintf("total-winnings-%v", limit)

	cachedUsers, err := db.cache.Get(ctx, cacheKey, new([]domain.User))
	if err == nil {
		users, ok := cachedUsers.([]domain.User)
		if !ok {
			return nil, fmt.Errorf("cannot cast interface to slice of user type")
		}

		return users, nil
	}

	users, err := db.query.GetTopUsers(ctx, limit)
	if err != nil {
		return nil, err
	}

	var mappedUsers []domain.User

	for _, user := range users {
		mappedUser := domain.User{
			ID:        user.UserID,
			TotalBets: user.TotalBets,
		}

		mappedUsers = append(mappedUsers, mappedUser)
	}

	err = db.cache.Set(ctx, cacheKey, mappedUsers, time.Minute)
	if err != nil {
		log.Println(err.Error())
	}

	return mappedUsers, nil
}

// GetAnomalousUsers fetches users with significantly higher betting activity than the average.
func (db MaybetsDB) GetAnomalousUsers(ctx context.Context) ([]domain.User, error) {
	_, span := tracer.Start(ctx, "GetTopUsers")
	defer span.End()

	cacheKey := "anomalous-users"

	cachedUsers, err := db.cache.Get(ctx, cacheKey, new([]domain.User))
	if err == nil {
		users, ok := cachedUsers.([]domain.User)
		if !ok {
			return nil, fmt.Errorf("cannot cast interface into users type")
		}

		return users, nil
	}

	users, err := db.query.GetAnomalousUsers(ctx)
	if err != nil {
		return nil, err
	}

	var mappedUsers []domain.User

	for _, user := range users {
		mappedUser := domain.User{
			ID:        user.UserID,
			TotalBets: user.TotalBets,
		}

		mappedUsers = append(mappedUsers, mappedUser)
	}

	err = db.cache.Set(ctx, cacheKey, mappedUsers, time.Minute)
	if err != nil {
		log.Println(err.Error())
	}

	return mappedUsers, nil
}
