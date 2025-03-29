package redis

import "go.uber.org/fx"

var RedisModule = fx.Module("liquor-redis-module", fx.Provide(NewRedisClient))
