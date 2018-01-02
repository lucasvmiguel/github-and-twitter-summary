package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasvmiguel/github-and-twitter-summary/github"
	"github.com/lucasvmiguel/github-and-twitter-summary/summarize"
	"github.com/lucasvmiguel/github-and-twitter-summary/twitter"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = viper.GetString("github-and-twitter-summary")
	app.Usage = viper.GetString("Command line to summarize responses from github api and twitter api")
	app.Version = viper.GetString("0.0.1")
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		viper.SetConfigFile(c.String("c"))
		err := viper.ReadInConfig()

		if err != nil {
			fmt.Printf("Error: Couldn't read config file: %s", err.Error())
			os.Exit(1)
		}

		githubClient := github.NewClient(github.Config{
			Token:   viper.GetString("github.token"),
			PerPage: viper.GetInt("github.per_page"),
		})

		twitterClient := twitter.NewClient(twitter.Config{
			AccessToken:       viper.GetString("twitter.access_token"),
			AccessTokenSecret: viper.GetString("twitter.access_token_secret"),
			ConsumerKey:       viper.GetString("twitter.consumer_key"),
			ConsumerSecret:    viper.GetString("twitter.consumer_secret"),
		})

		summarizeClient := summarize.NewClient(summarize.Config{
			GithubClient:  githubClient,
			TwitterClient: twitterClient,
		})

		summaries, err := summarizeClient.SummarizeByText(context.Background(), c.Args().First())
		if err != nil {
			return err
		}

		summarize.BeautifulPrint(summaries)

		return nil
	}

	app.Run(os.Args)
}
