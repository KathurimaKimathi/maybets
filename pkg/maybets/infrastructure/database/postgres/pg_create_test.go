package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/enums"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
	cacheMock "github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/cache/mock"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	gormMock "github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm/mock"
	"github.com/brianvoe/gofakeit"
)

func TestMaybetsDB_StoreBetData(t *testing.T) {
	type args struct {
		ctx  context.Context
		bets []*domain.Bet
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: store bets in db",
			args: args{
				ctx: context.Background(),
				bets: []*domain.Bet{
					{BetID: gofakeit.UUID(), UserID: gofakeit.UUID(), Amount: 198, Odds: 3.2, Outcome: enums.Win, Timestamp: time.Now()},
				},
			},
			wantErr: false,
		},
		{
			name: "sad: unable to store bets in db",
			args: args{
				ctx: context.Background(),
				bets: []*domain.Bet{
					{BetID: gofakeit.UUID(), UserID: gofakeit.UUID(), Amount: 198, Odds: 3.2, Outcome: enums.Win, Timestamp: time.Now()},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeGorm := gormMock.NewGormMock()
			fakeCache := cacheMock.NewStoreCacheMock()

			db := NewMaybetsDB(fakeCache, fakeGorm, fakeGorm)

			if tt.name == "sad: unable to store bets in db" {
				fakeGorm.MockStoreBetDataFn = func(_ context.Context, _ []gorm.Bet) error {
					return fmt.Errorf("error")
				}
			}

			if err := db.StoreBetData(tt.args.ctx, tt.args.bets); (err != nil) != tt.wantErr {
				t.Errorf("MaybetsDB.StoreBetData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
