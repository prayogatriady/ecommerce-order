package database

import (
	configM "github.com/prayogatriady/ecommerce-module/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     configM.String("redis.host", ""),
		Password: configM.String("redis.password", ""),
		DB:       0,
	})

	return client
}
