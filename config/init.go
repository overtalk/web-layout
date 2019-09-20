package config

import (
	"web-layout/config/env"
	"web-layout/config/gitlab"
)

// Mode defines the server running mode
type Mode int

const (
	DEV = iota
	PROD
)

// NewConfig is the resolver of config struct
// input arg `config` should be pointer of config Struct
func NewConfig(mode Mode, config interface{}) error {
	var err error
	switch mode {
	case DEV:
		err = gitlab.Fetch(config)
	default:
		err = env.Fetch(config)
	}

	return err
}
