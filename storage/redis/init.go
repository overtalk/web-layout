package redis

import (
	"github.com/go-redis/redis"

	redisPool "web-layout/utils/redis"
)

type Redis struct {
	config redisPool.Config
	client redis.Cmdable
}

func NewRedis(c redisPool.Config) (*Redis, error) {
	client, err := c.Connect()
	if err != nil {
		return nil, err
	}

	return &Redis{
		config: c,
		client: client,
	}, nil
}
