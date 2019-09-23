package mem

import (
	"web-layout/utils/memcache"
)

type Memory struct {
	client *concurrentcache.ConcurrentCache
}

func NewMemory() (*Memory, error) {
	client, err := concurrentcache.NewConcurrentCache(256, 10240)
	if err != nil {
		return nil, err
	}

	return &Memory{
		client: client,
	}, nil
}
