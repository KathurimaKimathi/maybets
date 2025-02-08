package gorm_test

import (
	"context"
	"testing"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	"github.com/brianvoe/gofakeit"
)

func TestDBInstance_StoreBetData(t *testing.T) {
	type args struct {
		ctx context.Context
		bet []gorm.Bet
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
				bet: []gorm.Bet{
					{BetID: gofakeit.UUID(), UserID: userID, Amount: 100, Odds: 2.78, Outcome: "win", Timestamp: time.Now()},
					{BetID: gofakeit.UUID(), UserID: userID2, Amount: 59, Odds: 1.78, Outcome: "lose", Timestamp: time.Now()},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad: unable to store bets in db",
			args: args{
				ctx: context.Background(),
				bet: []gorm.Bet{
					{BetID: gofakeit.UUID(), UserID: userID, Amount: 100, Odds: 2.78, Outcome: "win", Timestamp: time.Now()},
					{BetID: bet1UserID, UserID: userID2, Amount: 59, Odds: 1.78, Outcome: "lose", Timestamp: time.Now()},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testingDB.StoreBetData(tt.args.ctx, tt.args.bet); (err != nil) != tt.wantErr {
				t.Errorf("DBInstance.StoreBetData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
