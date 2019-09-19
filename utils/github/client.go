package github

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Config struct {
	Token string `env:"GITHUB_TOKEN"`
	Owner string `env:"GITHUB_OWNER"`
	Repo  string `env:"GITHUB_REPO"`
	Ref   string `env:"GITHUB_REF" envDefault:"master"`
}

type client struct {
	cfg Config
	git *github.Client
}

func NewClient(c Config) *client {
	cli := github.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: c.Token},
			)),
	)
	return &client{git: cli, cfg: c}
}

func (c *client) Client() (*github.Client, error) {
	if c.git != nil {
		return nil, errors.New("github client is nil")
	}
	return c.git, nil
}
