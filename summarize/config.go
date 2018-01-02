package summarize

import (
	"github.com/lucasvmiguel/github-and-twitter-summary/github"
	"github.com/lucasvmiguel/github-and-twitter-summary/twitter"
)

// Config struct to config package
type Config struct {
	GithubClient  github.Client
	TwitterClient twitter.Client
}
