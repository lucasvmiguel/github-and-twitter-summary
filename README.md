# github-and-twitter-summary
[![Build Status](https://travis-ci.org/lucasvmiguel/github-and-twitter-summary.svg?branch=master)](https://travis-ci.org/lucasvmiguel/github-and-twitter-summary)
[![GoDoc](https://godoc.org/github.com/lucasvmiguel/github-and-twitter-summary?status.svg)](https://godoc.org/github.com/lucasvmiguel/github-and-twitter-summary)

## Overview

A command line that get repositories by name and show tweets related.

## Architeture Explanation

[Here](https://medium.com/@LucasVieiraDev/dependencies-in-golang-projects-f46a11fef832)

## Installation

Make sure you have a working Go environment. [See
the install instructions for Go](http://golang.org/doc/install.html).

To install github-and-twitter-summary, simply run:
```
$ go get github.com/lucasvmiguel/github-and-twitter-summary
```

## Configuration

```
[github]
token="ACCESS-TOKEN"
per_page=NUMBER-REPOSITORIES-PER-PAGE

[twitter]
access_token="ACCESS-TOKEN"
access_token_secret="ACCESS-TOKEN-SECRET"
consumer_key="CONSUMER-KEY"
consumer_secret="CONSUMER-SECRET"
```

## Usage

command:
```
$ github-and-twitter-summary -c "<PATH-TO-CONFIG-FILE>" "<TEXT-TO-SEARCH>"
```

## Example

config file:
```
[github]
token="abcdefghi"
per_page=10

[twitter]
access_token="abcdefghi"
access_token_secret="abcdefghi"
consumer_key="abcdefghi"
consumer_secret="abcdefghi"
```

command:
```
$ github-and-twitter-summary -c "/home/myname/config.toml" "football"
```

## Documentation

[Here](https://godoc.org/github.com/lucasvmiguel/github-and-twitter-summary)