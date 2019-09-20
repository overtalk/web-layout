package env

import "github.com/caarlos0/env"

// Fetch gets all config items form env
func Fetch(config interface{}) error {
	return env.Parse(config)
}
