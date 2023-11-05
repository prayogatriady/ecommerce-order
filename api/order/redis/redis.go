package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Create(ctx context.Context, key string, val any) error
	Find(ctx context.Context, key string) (string, error)
}

type redisRepository struct {
	*redis.Client
	duration time.Duration
}

func NewRedisRepository(rdb *redis.Client) RedisRepository {
	return &redisRepository{
		Client:   rdb,
		duration: 5 * time.Minute,
	}
}

func (r *redisRepository) Create(ctx context.Context, key string, val any) error {
	return r.Set(ctx, key, val, r.duration).Err()
}

func (r *redisRepository) Find(ctx context.Context, key string) (string, error) {
	return r.Get(ctx, key).Result()
}
