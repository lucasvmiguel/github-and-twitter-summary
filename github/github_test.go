package github

import "testing"

func TestNewClient(t *testing.T) {
	cfg := Config{PerPage: 12, Token: "abc"}

	client := NewClient(cfg)

	if client.Api == nil {
		t.Fatal("Expected not nil but got nil")
	}

	if client.config.PerPage != 12 {
		t.Fatalf("Expected 12 but got %d", client.config.PerPage)
	}

	if client.config.Token != "abc" {
		t.Fatalf("Expected abc but got %s", client.config.Token)
	}
}
