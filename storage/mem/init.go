package mem

import (
	"web-layout/utils/memcache"
)

type memCache struct {
	client *concurrentcache.ConcurrentCache
}

func Mem() (*memCache, error) {
	client, err := concurrentcache.NewConcurrentCache(256, 10240)
	if err != nil {
		return nil, err
	}

	return &memCache{
		client: client,
	}, nil
}
