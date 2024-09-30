package database

import (
	"context"
	"fmt"
	"http-server/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	db *redis.Client
}

var (
	redisInstance *cache
	redisOnce     sync.Once
)

func ConnectToCache(ctx context.Context, cfg *config.Config) *cache {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			Password: "",
			DB:       0,
		})

		redisInstance = &cache{client}
	})

	return redisInstance
}

func (c *cache) Ping(ctx context.Context) (string, error) {
	return c.db.Ping(ctx).Result()
}

func (c *cache) Close() {
	c.db.Close()
}

func GetCache() *redis.Client {
	return redisInstance.db
}
