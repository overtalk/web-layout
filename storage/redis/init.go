package redis

import (
	"github.com/go-redis/redis"

	redisPool "web-layout/utils/redis"
)

type redisCache struct {
	config redisPool.Config
	client redis.Cmdable
}

func Redis(c redisPool.Config) (*redisCache, error) {
	client, err := c.Connect()
	if err != nil {
		return nil, err
	}

	return &redisCache{
		config: c,
		client: client,
	}, nil
}
