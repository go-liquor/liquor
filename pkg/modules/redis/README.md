# Redis Module

## Contents
- [Enable](#enable)
- [Available Operations](#available-operations)
  - [Key-Value Operations](#key-value-operations)
  - [Hash Operations](#hash-operations)
  - [List Operations](#list-operations)
- [Usage Example](#usage-example)
- [In-Memory Implementation](#in-memory-implementation)
- [Testing](#testing)


In `cmd/app/main.go` add module:

```go
package main

import (
    "github.com/go-liquor/liquor/pkg/modules/redis"
)

func main() {
    app.NewApp(
        redis.RedisModule,
    )
}
```

Add this in your `config.yaml`:

```yaml
redis:
  addr: "<REDIS_ADDR>" #eg localhost:6379
  password: ""  # optional
```

## Available Operations

### Key-Value Operations

```go
type RedisClient interface {
    // Store a key with optional expiration
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    
    // Retrieve a key's value
    Get(ctx context.Context, key string) (string, error)
    
    // Delete one or more keys
    Delete(ctx context.Context, keys ...string) error
    
    // Check if keys exist
    Exists(ctx context.Context, keys ...string) (bool, error)
    
    // Set key expiration
    Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
    
    // Increment a counter
    Incr(ctx context.Context, key string) (int64, error)
}
```

### Hash Operations

```go
type RedisClient interface {
    // Set hash fields
    HSet(ctx context.Context, key string, values ...interface{}) error
    
    // Get a hash field
    HGet(ctx context.Context, key, field string) (string, error)
    
    // Get all hash fields
    HGetAll(ctx context.Context, key string) (map[string]string, error)
    
    // Delete hash fields
    HDel(ctx context.Context, key string, fields ...string) error
}
```

### List Operations

```go
type RedisClient interface {
    // Add to list head
    LPush(ctx context.Context, key string, values ...interface{}) error
    
    // Remove from list head
    LPop(ctx context.Context, key string) (string, error)
    
    // Add to list tail
    RPush(ctx context.Context, key string, values ...interface{}) error
    
    // Remove from list tail
    RPop(ctx context.Context, key string) (string, error)
}
```

## Usage Example

```go
type Service struct {
    redis RedisClient
}

func NewService(redis RedisClient) *Service {
    return &Service{
        redis: redis,
    }
}

func (s *Service) CacheData(ctx context.Context) error {
    // Store data with 1 hour expiration
    err := s.redis.Set(ctx, "my-key", "my-value", time.Hour)
    if err != nil {
        return err
    }

    // Store hash data
    err = s.redis.HSet(ctx, "user:123", "name", "John", "age", "30")
    if err != nil {
        return err
    }

    // Add to list
    err = s.redis.LPush(ctx, "notifications", "new-message")
    if err != nil {
        return err
    }

    return nil
}
```

## In-Memory Implementation

For testing or development purposes, you can use the in-memory implementation:

```go
func main() {
    // Use in-memory implementation
    redis := redis.NewInMemoryRedis()
    
    service := NewService(redis)
    // ... use service
}
```

## Testing

```go
func TestYourService(t *testing.T) {
    // Using in-memory implementation
    redis := redis.NewInMemoryRedis()
    service := NewService(redis)

    // Test your service
    err := service.CacheData(context.Background())
    require.NoError(t, err)

    // Verify data
    value, err := redis.Get(context.Background(), "my-key")
    require.NoError(t, err)
    assert.Equal(t, "my-value", value)
}
```