package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/go-redis/redis"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/cache")

// StoreCache provides a higher-level abstraction for interacting with the cache.
// It uses a StoreCacheService implementation to perform the actual cache operations.
type StoreCache struct {
	storer *redis.Client
}

// NewStoreCache creates a new StoreCache instance.
// It takes a StoreCache implementation as an argument, which will be used for cache operations.
// This allows for flexible cache service implementations to be passed in.
func NewStoreCache(storer *redis.Client) *StoreCache {
	return &StoreCache{
		storer: storer,
	}
}

// Get fetches a value from the Redis cache by its key and returns it as the specified type.
func (cs *StoreCache) Get(ctx context.Context, key string, valueType interface{}) (interface{}, error) {
	_, span := tracer.Start(ctx, "Get")
	defer span.End()

	val, err := cs.storer.Get(key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("value with key, %s, not found", key)
		}

		return nil, fmt.Errorf("failed to get value from Redis: %w", err)
	}

	// Check if valueType is a non-nil pointer
	valueTypeVal := reflect.ValueOf(valueType)
	if valueTypeVal.Kind() != reflect.Ptr || valueTypeVal.IsNil() {
		return nil, fmt.Errorf("valueType must be a non-nil pointer")
	}

	// Convert string value from Redis to the specified type
	// This assumes the value stored in Redis is JSON encoded
	err = json.Unmarshal([]byte(val), valueType)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Redis value: %w", err)
	}

	return reflect.ValueOf(valueType).Elem().Interface(), nil
}

// Set stores a value in the cache with the given key and expiration duration.
func (cs *StoreCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	_, span := tracer.Start(ctx, "Set")
	defer span.End()

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	err = cs.storer.Set(key, data, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %w", err)
	}

	return nil
}
