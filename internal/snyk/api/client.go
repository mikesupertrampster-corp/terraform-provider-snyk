package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	ctx        context.Context
	hostUrl    string
	httpClient *http.Client
	orgId      string
	userAgent  string
}

var ErrInvalidAuthn = errors.New("credentials not valid")
var ErrInvalidAuthz = errors.New("credentials not authorized to access resource")
var ErrNotFound = errors.New("requested resource not found")

func NewClient(hostUrl string, apiKey string, orgId string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}

	return &Client{
		apiKey:     apiKey,
		hostUrl:    hostUrl,
		httpClient: client,
		orgId:      orgId,
		userAgent:  "terraform-provider-snyk",
	}
}

func (c *Client) doRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.hostUrl, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, ErrInvalidAuthn
	} else if resp.StatusCode == 403 {
		return nil, ErrInvalidAuthz
	} else if resp.StatusCode == 404 {
		return nil, ErrNotFound
	} else if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Validate(ctx context.Context) error {
	_, err := c.doRequest(ctx, http.MethodGet, "user/me", []byte{})
	if err != nil {
		return err
	}

	return nil
}
