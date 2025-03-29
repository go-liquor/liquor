package redis

import (
	"context"
	"time"

	"github.com/go-liquor/liquor/v2/config"
	goredis "github.com/redis/go-redis/v9"
)

// RedisClient defines the interface for Redis operations
type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (bool, error)
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Incr(ctx context.Context, key string) (int64, error)
	HSet(ctx context.Context, key string, values ...interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	LPush(ctx context.Context, key string, values ...interface{}) error
	LPop(ctx context.Context, key string) (string, error)
	RPush(ctx context.Context, key string, values ...interface{}) error
	RPop(ctx context.Context, key string) (string, error)
}

type redisClient struct {
	client *goredis.Client
}

// NewRedisClient creates a new Redis client using the provided configuration.
//
// Parameters:
//   - cfg: Configuration object containing Redis connection details
//
// Returns:
//   - RedisClient: Interface for Redis operations
func NewRedisClient(cfg *config.Config) RedisClient {
	clt := goredis.NewClient(&goredis.Options{
		Addr:     cfg.GetString("redis.addr"),
		Password: cfg.GetString("redis.password"),
	})
	return &redisClient{
		client: clt,
	}
}

// Set stores a key-value pair with optional expiration.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Key to store
//   - value: Value to store
//   - expiration: Time until the key expires (0 for no expiration)
//
// Returns:
//   - error: nil if successful, error otherwise
func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by its key.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Key to retrieve
//
// Returns:
//   - string: Retrieved value
//   - error: nil if successful, error otherwise
func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete removes one or more keys.
//
// Parameters:
//   - ctx: Context for the operation
//   - keys: Keys to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (r *redisClient) Delete(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// Exists checks if one or more keys exist.
//
// Parameters:
//   - ctx: Context for the operation
//   - keys: Keys to check
//
// Returns:
//   - bool: true if all keys exist, false otherwise
//   - error: nil if successful, error otherwise
func (r *redisClient) Exists(ctx context.Context, keys ...string) (bool, error) {
	result, err := r.client.Exists(ctx, keys...).Result()
	return result > 0, err
}

// Expire sets an expiration time on a key.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Key to set expiration on
//   - expiration: Time until the key expires
//
// Returns:
//   - bool: true if the timeout was set, false otherwise
//   - error: nil if successful, error otherwise
func (r *redisClient) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return r.client.Expire(ctx, key, expiration).Result()
}

// Incr increments a number stored at key.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Key to increment
//
// Returns:
//   - int64: New value after increment
//   - error: nil if successful, error otherwise
func (r *redisClient) Incr(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

// HSet sets multiple hash fields to multiple values.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Hash key
//   - values: Field-value pairs to set
//
// Returns:
//   - error: nil if successful, error otherwise
func (r *redisClient) HSet(ctx context.Context, key string, values ...interface{}) error {
	return r.client.HSet(ctx, key, values...).Err()
}

// HGet retrieves the value of a hash field.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Hash key
//   - field: Field to get
//
// Returns:
//   - string: Value of the field
//   - error: nil if successful, error otherwise
func (r *redisClient) HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

// HGetAll retrieves all fields and values of a hash.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Hash key
//
// Returns:
//   - map[string]string: All fields and their values
//   - error: nil if successful, error otherwise
func (r *redisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

// HDel removes one or more hash fields.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: Hash key
//   - fields: Fields to remove
//
// Returns:
//   - error: nil if successful, error otherwise
func (r *redisClient) HDel(ctx context.Context, key string, fields ...string) error {
	return r.client.HDel(ctx, key, fields...).Err()
}

// LPush inserts values at the head of a list.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: List key
//   - values: Values to push
//
// Returns:
//   - error: nil if successful, error otherwise
func (r *redisClient) LPush(ctx context.Context, key string, values ...interface{}) error {
	return r.client.LPush(ctx, key, values...).Err()
}

// LPop removes and returns the first element of a list.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: List key
//
// Returns:
//   - string: The popped element
//   - error: nil if successful, error otherwise
func (r *redisClient) LPop(ctx context.Context, key string) (string, error) {
	return r.client.LPop(ctx, key).Result()
}

// RPush inserts values at the tail of a list.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: List key
//   - values: Values to push
//
// Returns:
//   - error: nil if successful, error otherwise
func (r *redisClient) RPush(ctx context.Context, key string, values ...interface{}) error {
	return r.client.RPush(ctx, key, values...).Err()
}

// RPop removes and returns the last element of a list.
//
// Parameters:
//   - ctx: Context for the operation
//   - key: List key
//
// Returns:
//   - string: The popped element
//   - error: nil if successful, error otherwise
func (r *redisClient) RPop(ctx context.Context, key string) (string, error) {
	return r.client.RPop(ctx, key).Result()
}
