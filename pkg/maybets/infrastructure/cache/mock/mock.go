package mock

import (
	"context"
	"time"
)

// StoreCacheMock mocks caching implementations
type StoreCacheMock struct {
	MockGetFn func(ctx context.Context, key string, valueType interface{}) (interface{}, error)
	MockSetFn func(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

// NewStoreCacheMock initializes our client mocks
func NewStoreCacheMock() *StoreCacheMock {
	return &StoreCacheMock{
		MockGetFn: func(ctx context.Context, key string, valueType interface{}) (interface{}, error) { //nolint:all
			return "", nil
		},
		MockSetFn: func(ctx context.Context, key string, value interface{}, expiration time.Duration) error { //nolint:all
			return nil
		},
	}
}

// Get mocks the implementation of getting a value from cache
func (c StoreCacheMock) Get(ctx context.Context, key string, valueType interface{}) (interface{}, error) {
	return c.MockGetFn(ctx, key, valueType)
}

// Set mocks the implementation of setting a value to the cache store
func (c StoreCacheMock) Set(ctx context.Context, key string, valueType interface{}, expiration time.Duration) error {
	return c.MockSetFn(ctx, key, valueType, expiration)
}
