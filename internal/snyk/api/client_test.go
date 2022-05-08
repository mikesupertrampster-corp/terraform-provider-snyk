package api

import (
	"context"
	"testing"
	"time"
)

func TestClientInvalidation(t *testing.T) {
	var tests = []struct {
		hostUrl string
		apiKey  string
		want    error
	}{
		{
			"https://snyk.io/api/v1",
			"00000000-0000-0000-0000-000000000000",
			ErrInvalidAuthn,
		},
		{
			"https://snyk.io/no/such/path",
			"00000000-0000-0000-0000-000000000000",
			ErrNotFound,
		},
	}

	for _, tt := range tests {
		c := NewClient(
			tt.hostUrl,
			tt.apiKey,
			"00000000-0000-0000-0000-000000000000",
			10*time.Second,
		)

		err := c.Validate(context.Background())
		if err != tt.want {
			t.Fatal(err)
		}
	}
}
