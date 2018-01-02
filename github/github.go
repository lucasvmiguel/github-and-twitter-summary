package github

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Client struct to use package
type Client struct {
	config Config
	Api    *github.Client
}

// NewClient receives the configuration to access the API and returns the client that make the HTTP requests
func NewClient(config Config) Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return Client{config: config, Api: github.NewClient(tc)}
}

// GetRepositoriesByText returns repositories by text passed as parameter
func (c *Client) GetRepositoriesByText(ctx context.Context, text string) ([]github.Repository, error) {
	opt := &github.SearchOptions{ListOptions: github.ListOptions{PerPage: c.config.PerPage}}
	response, _, err := c.Api.Search.Repositories(ctx, text, opt)
	if err != nil {
		return response.Repositories, errors.Wrap(err, "couldn't find repositories by text")
	}
	return response.Repositories, nil
}
