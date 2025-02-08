package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
	cacheMock "github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/cache/mock"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	gormMock "github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm/mock"
	"github.com/google/uuid"
)

func TestMaybetsDB_GetTotalBets(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: get bets from db",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: false,
		},
		{
			name: "success: get bets from cache",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: false,
		},
		{
			name: "fail: invalid type in cache",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name: "fail: fail to get from db",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name: "success: get bets but not set in cache",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeGorm := gormMock.NewGormMock()
			fakeCache := cacheMock.NewStoreCacheMock()

			db := NewMaybetsDB(fakeCache, fakeGorm, fakeGorm)

			if tt.name == "success: get bets from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}
			}

			if tt.name == "success: get bets from cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					total := int64(5)

					return &total, nil
				}
			}

			if tt.name == "success: get bets but not set in cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeCache.MockSetFn = func(_ context.Context, _ string, _ interface{}, _ time.Duration) error {
					return fmt.Errorf("error")
				}
			}

			if tt.name == "fail: fail to get from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeGorm.MockGetTotalBetsFn = func(_ context.Context, _ string) (int64, error) {
					return 0, fmt.Errorf("error")
				}
			}

			_, err := db.GetTotalBets(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaybetsDB.GetTotalBets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMaybetsDB_GetTotalWinnings(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: get bets from db",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: false,
		},
		{
			name: "success: get bets from cache",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: false,
		},
		{
			name: "fail: invalid type in cache",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name: "fail: fail to get from db",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name: "success: get bets but not set in cache",
			args: args{
				ctx:    context.Background(),
				userID: uuid.NewString(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeGorm := gormMock.NewGormMock()
			fakeCache := cacheMock.NewStoreCacheMock()

			db := NewMaybetsDB(fakeCache, fakeGorm, fakeGorm)

			if tt.name == "success: get bets from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}
			}

			if tt.name == "success: get bets from cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					total := float64(5)

					return &total, nil
				}
			}

			if tt.name == "success: get bets but not set in cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeCache.MockSetFn = func(_ context.Context, _ string, _ interface{}, _ time.Duration) error {
					return fmt.Errorf("error")
				}
			}

			if tt.name == "fail: fail to get from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeGorm.MockGetTotalWinningsFn = func(_ context.Context, _ string) (float64, error) {
					return 0, fmt.Errorf("error")
				}
			}

			_, err := db.GetTotalWinnings(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaybetsDB.GetTotalWinnings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMaybetsDB_GetTopUsers(t *testing.T) {
	type args struct {
		ctx   context.Context
		limit int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: get bets from db",
			args: args{
				ctx:   context.Background(),
				limit: 5,
			},
			wantErr: false,
		},
		{
			name: "success: get bets from cache",
			args: args{
				ctx:   context.Background(),
				limit: 5,
			},
			wantErr: false,
		},
		{
			name: "fail: invalid type in cache",
			args: args{
				ctx:   context.Background(),
				limit: 5,
			},
			wantErr: true,
		},
		{
			name: "fail: fail to get from db",
			args: args{
				ctx:   context.Background(),
				limit: 5,
			},
			wantErr: true,
		},
		{
			name: "success: get bets but not set in cache",
			args: args{
				ctx:   context.Background(),
				limit: 5,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeGorm := gormMock.NewGormMock()
			fakeCache := cacheMock.NewStoreCacheMock()

			db := NewMaybetsDB(fakeCache, fakeGorm, fakeGorm)

			if tt.name == "success: get bets from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}
			}

			if tt.name == "success: get bets from cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return []domain.User{
						{
							ID:        uuid.NewString(),
							TotalBets: 5,
						},
					}, nil
				}
			}

			if tt.name == "success: get bets but not set in cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeCache.MockSetFn = func(_ context.Context, _ string, _ interface{}, _ time.Duration) error {
					return fmt.Errorf("error")
				}
			}

			if tt.name == "fail: fail to get from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeGorm.MockGetTopUsersFn = func(_ context.Context, _ int) ([]gorm.User, error) {
					return nil, fmt.Errorf("error")
				}
			}

			_, err := db.GetTopUsers(tt.args.ctx, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaybetsDB.GetTopUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMaybetsDB_GetAnomalousUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: get bets from db",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "success: get bets from cache",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "fail: invalid type in cache",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "fail: fail to get from db",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "success: get bets but not set in cache",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeGorm := gormMock.NewGormMock()
			fakeCache := cacheMock.NewStoreCacheMock()

			db := NewMaybetsDB(fakeCache, fakeGorm, fakeGorm)

			if tt.name == "success: get bets from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}
			}

			if tt.name == "success: get bets from cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return []domain.User{
						{
							ID:        uuid.NewString(),
							TotalBets: 5,
						},
					}, nil
				}
			}

			if tt.name == "success: get bets but not set in cache" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeCache.MockSetFn = func(_ context.Context, _ string, _ interface{}, _ time.Duration) error {
					return fmt.Errorf("error")
				}
			}

			if tt.name == "fail: fail to get from db" {
				fakeCache.MockGetFn = func(_ context.Context, _ string, _ interface{}) (interface{}, error) {
					return nil, fmt.Errorf("error")
				}

				fakeGorm.MockGetAnomalousUsersFn = func(_ context.Context) ([]gorm.User, error) {
					return nil, fmt.Errorf("error")
				}
			}

			_, err := db.GetAnomalousUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaybetsDB.GetTopUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
