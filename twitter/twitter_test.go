package twitter

import "testing"

func TestNewClient(t *testing.T) {
	cfg := Config{AccessToken: "abc", AccessTokenSecret: "def", ConsumerKey: "ghi", ConsumerSecret: "jkl"}

	client := NewClient(cfg)

	if client.Api == nil {
		t.Fatal("Expected not nil but got nil")
	}

	if client.config.AccessToken != "abc" {
		t.Fatalf("Expected abc but got %s", client.config.AccessToken)
	}

	if client.config.AccessTokenSecret != "def" {
		t.Fatalf("Expected abc but got %s", client.config.AccessTokenSecret)
	}

	if client.config.ConsumerKey != "ghi" {
		t.Fatalf("Expected abc but got %s", client.config.ConsumerKey)
	}

	if client.config.ConsumerSecret != "jkl" {
		t.Fatalf("Expected abc but got %s", client.config.ConsumerSecret)
	}
}
