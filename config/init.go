package config

import (
	"web-layout/config/env"
	"web-layout/config/gitlab"
)

// Mode defines the server running mode
// Mode defines the config source
type Mode int

// fetcher defines the uniform handler to get config from source
type fetcher func(interface{}) error

const (
	// DEV mode get config from gitlab repo source
	DEV = iota
	// PROD mode get config from env variables
	PROD
)

var fetchers map[Mode]fetcher

func init() {
	// collocate server running mode & config source
	fetchers = map[Mode]fetcher{
		DEV:  gitlab.Fetch,
		PROD: env.Fetch,
	}
}

// NewConfig is the resolver of config struct
// input arg `config` should be pointer of config Struct
func NewConfig(mode Mode, config interface{}) error {
	return fetchers[mode](config)
}
