package summarize

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/lucasvmiguel/github-and-twitter-summary/github"
	"github.com/lucasvmiguel/github-and-twitter-summary/twitter"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// Client struct to use package
type Client struct {
	githubClient  github.Client
	twitterClient twitter.Client
}

// Summary struct to show repository name and tweets
type Summary struct {
	RepositoryName string
	Tweets         []string
}

// NewClient receives the configuration to access the APIs
func NewClient(config Config) Client {
	return Client{
		githubClient:  config.GithubClient,
		twitterClient: config.TwitterClient,
	}
}

// BeautifulPrint print in terminal summaries passed as parameter
func BeautifulPrint(summaries []*Summary) {
	for _, summary := range summaries {
		table := tablewriter.NewWriter(os.Stdout)

		if len(summary.Tweets) < 1 {
			table.SetHeader([]string{summary.RepositoryName})
			table.Append([]string{"there aren't tweets for this repository"})

			table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor})
		} else {
			table.SetHeader([]string{"", summary.RepositoryName, ""})

			table.SetHeaderColor(
				tablewriter.Colors{},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
				tablewriter.Colors{})
		}

		for i, tweet := range summary.Tweets {
			table.Append([]string{strconv.Itoa(i + 1), tweet})
			table.Append([]string{""})
		}

		table.Render()
	}
}

// SummarizeByText return summaries by text passed as parameter
func (c *Client) SummarizeByText(ctx context.Context, text string) ([]*Summary, error) {
	repositories, err := c.githubClient.GetRepositoriesByText(ctx, text)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find repositories by text")
	}

	var wg sync.WaitGroup
	summaries := []*Summary{}
	wg.Add(len(repositories))

	for _, repository := range repositories {
		summary := &Summary{}
		summaries = append(summaries, summary)
		summary.RepositoryName = *repository.FullName

		go c.fillSummaryWithTweets(&wg, summary)
	}

	wg.Wait()

	return summaries, nil
}

func (c *Client) fillSummaryWithTweets(wg *sync.WaitGroup, summary *Summary) {
	tweets, err := c.twitterClient.GetTweetsByText(summary.RepositoryName)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, tweet := range tweets {
		summary.Tweets = append(summary.Tweets, tweet.Text)
	}

	wg.Done()
}
