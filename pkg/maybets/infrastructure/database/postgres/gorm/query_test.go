package gorm_test

import (
	"context"
	"testing"
)

func TestDBInstance_GetTotalBets(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success: user with 4 bets",
			args: args{
				ctx:    context.Background(),
				userID: userID,
			},
			want:    int64(4),
			wantErr: false,
		},
		{
			name: "success: user with 5 bets",
			args: args{
				ctx:    context.Background(),
				userID: userID2,
			},
			want:    int64(7),
			wantErr: false,
		},
		{
			name: "success: user with 2 bets",
			args: args{
				ctx:    context.Background(),
				userID: userID3,
			},
			want:    int64(2),
			wantErr: false,
		},
		{
			name: "success: no user with bets",
			args: args{
				ctx:    context.Background(),
				userID: "foo",
			},
			want:    0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testingDB.GetTotalBets(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBInstance.GetTotalBets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("DBInstance.GetTotalBets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBInstance_GetTotalWinnings(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success: user with 2 wins",
			args: args{
				ctx:    context.Background(),
				userID: userID,
			},
			want:    float64(169.00),
			wantErr: false,
		},
		{
			name: "success: user with no win",
			args: args{
				ctx:    context.Background(),
				userID: userID6,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testingDB.GetTotalWinnings(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBInstance.GetTotalWinnings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("DBInstance.GetTotalWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBInstance_GetTopUsers(t *testing.T) {
	type args struct {
		ctx   context.Context
		limit int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success: get 5",
			args: args{
				ctx:   context.Background(),
				limit: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "success: get all",
			args: args{
				ctx:   context.Background(),
				limit: 10,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "success: get 1",
			args: args{
				ctx:   context.Background(),
				limit: 1,
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testingDB.GetTopUsers(tt.args.ctx, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBInstance.GetTopUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != len(got) {
				t.Errorf("DBInstance.GetTopUsers() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestDBInstance_GetAnomalousUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success: get the one user who is anomalous",
			args: args{
				ctx: context.Background(),
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testingDB.GetAnomalousUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBInstance.GetAnomalousUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != len(got) {
				t.Errorf("DBInstance.GetAnomalousUsers() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
