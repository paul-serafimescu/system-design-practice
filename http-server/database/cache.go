package database

import (
	"context"
	"fmt"
	"http-server/config"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type cache struct {
	db               *redis.Client
	expirationPubSub *redis.PubSub

	OnWebsocketExpiration func(serviceId string) error
	OnError               func(err error)
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

		expirationPubSub := client.PSubscribe(ctx, "__keyevent@0__:expired")

		redisInstance = &cache{db: client, expirationPubSub: expirationPubSub}
	})

	return redisInstance
}

func (c *cache) Ping(ctx context.Context) (string, error) {
	return c.db.Ping(ctx).Result()
}

func (c *cache) HandleKeyExpiration() {
	for msg := range c.expirationPubSub.Channel() {
		if strings.HasPrefix(msg.Payload, "websocket") {
			serviceId, _ := strings.CutPrefix(msg.Payload, "websocket:")
			err := c.OnWebsocketExpiration(serviceId)

			if err != nil {
				c.OnError(err)
			}
		} else {
			log.Error().Msgf("unknown key type: %s", msg.Payload)
		}
	}
}

func (c *cache) Close() {
	c.db.Close()
}

func GetCache() *redis.Client {
	return redisInstance.db
}
