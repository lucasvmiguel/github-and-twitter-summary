package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/pkg/errors"
)

// Client struct to use package
type Client struct {
	config Config
	Api    *twitter.Client
}

// NewClient receives the configuration to access the API and returns the client that make the HTTP requests
func NewClient(config Config) Client {
	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessTokenSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	return Client{config: config, Api: twitter.NewClient(httpClient)}
}

// GetTweetsByText returns tweets by text passed as parameter
func (c *Client) GetTweetsByText(text string) ([]twitter.Tweet, error) {
	search, _, err := c.Api.Search.Tweets(&twitter.SearchTweetParams{Query: text})
	if err != nil {
		return search.Statuses, errors.Wrap(err, "couldn't find tweets by text")
	}
	return search.Statuses, nil
}
