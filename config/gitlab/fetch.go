package gitlab

import (
	"encoding/json"

	"github.com/caarlos0/env"

	. "web-layout/utils/gitlab"
)

// Fetch gets all config items from gitlab repo
func Fetch(c interface{}) error {
	gitlabConfig := Config{}
	if err := env.Parse(&gitlabConfig); err != nil {
		return err
	}

	fetcher := gitlabConfig.NewClient()

	data, err := fetcher.Fetch(gitlabConfig.Path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}
