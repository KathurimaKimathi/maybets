package postgres

import (
	"context"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	"github.com/mitchellh/mapstructure"
)

func (db MaybetsDB) StoreBetData(ctx context.Context, bets []*domain.Bet) error {
	var betData []gorm.Bet

	if err := mapstructure.Decode(bets, &betData); err != nil {
		return err
	}

	err := db.create.StoreBetData(ctx, betData)
	if err != nil {
		return err
	}

	return nil
}
