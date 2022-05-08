package api

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetProjects(t *testing.T) {
	c := NewClient(
		"https://snyk.io/api/v1",
		os.Getenv("SNYK_API_KEY"),
		os.Getenv("SNYK_ORG_ID"),
		10*time.Second,
	)

	var tests = []struct {
		projectName string
		want        error
	}{
		{
			"mikesupertrampster-corp/blockchain:simple/go.mod",
			nil,
		},
		{
			"no/such/project",
			ErrNotFound,
		},
	}

	for _, tt := range tests {
		_, err := c.GetProject(context.Background(), tt.projectName)
		if err != tt.want {
			t.Fatal(err)
		}
	}
}

func TestGetAllProjectNames(t *testing.T) {
	c := NewClient(
		"https://snyk.io/api/v1",
		os.Getenv("SNYK_API_KEY"),
		os.Getenv("SNYK_ORG_ID"),
		10*time.Second,
	)

	names, err := c.GetAllProjectNames(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(names) == 0 {
		t.Fatal("expected non-zero list")
	}

	for _, element := range names {
		t.Log("[", element, "]")
	}
}
