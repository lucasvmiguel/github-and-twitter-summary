package summarize

import (
	"testing"

	"github.com/lucasvmiguel/github-and-twitter-summary/github"
	"github.com/lucasvmiguel/github-and-twitter-summary/twitter"
)

func TestNewClient(t *testing.T) {
	cfg := Config{GithubClient: github.Client{}, TwitterClient: twitter.Client{}}

	client := NewClient(cfg)

	if client.githubClient.Api != nil {
		t.Fatal("Expected not nil but got nil")
	}

	if client.twitterClient.Api != nil {
		t.Fatal("Expected not nil but got nil")
	}
}
