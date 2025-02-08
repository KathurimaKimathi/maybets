package cache

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestSNewStoreCache_Set(t *testing.T) {
	redisURL := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		t.Errorf("error parsing redis: %v", err)
		return
	}

	c := redis.NewClient(opt)

	type FooType string

	type args struct {
		ctx        context.Context
		key        string
		value      interface{}
		expiration time.Duration
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: set cache key",
			args: args{
				ctx:        context.Background(),
				key:        "foo",
				value:      "bar",
				expiration: 5 * time.Minute,
			},
		},
		{
			name: "success: set cache key",
			args: args{
				ctx:        context.Background(),
				key:        "bar",
				value:      FooType("foo"),
				expiration: 5 * time.Minute,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewStoreCache(c)
			if err := cs.Set(tt.args.ctx, tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("SNewStoreCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSNewStoreCache_Get(t *testing.T) {
	redisURL := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		t.Errorf("error parsing redis: %v", err)
		return
	}

	c := redis.NewClient(opt)

	type FooType string

	type args struct {
		ctx       context.Context
		key       string
		valueType interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "success: key found",
			args: args{
				ctx:       context.Background(),
				key:       "foo",
				valueType: new(FooType),
			},
			want:    FooType("bar"),
			wantErr: false,
		},
		{
			name: "fail: key not found",
			args: args{
				ctx:       context.Background(),
				key:       "haiko",
				valueType: new(FooType),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail: nil pointer",
			args: args{
				ctx:       context.Background(),
				key:       "foo",
				valueType: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewStoreCache(c)

			err := cs.Set(context.Background(), "foo", FooType("bar"), 10)
			if err != nil {
				t.Errorf("NewStoreCache.Set() got an error: %v", err)
				return
			}

			_, err = cs.Get(tt.args.ctx, tt.args.key, tt.args.valueType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStoreCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
