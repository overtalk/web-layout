package github

import (
	"encoding/json"

	"github.com/caarlos0/env"

	"web-layout/utils/github"
)

func Fetch(c interface{}) error {
	githubConfig := github.Config{}
	if err := env.Parse(&githubConfig); err != nil {
		return err
	}

	fetcher := githubConfig.NewClient()

	data, err := fetcher.Fetch()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}
