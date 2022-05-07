package snyk

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
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

func NewClient(ctx context.Context, hostUrl string, apiKey string, orgId string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}

	return &Client{
		apiKey:     apiKey,
		ctx:        ctx,
		hostUrl:    hostUrl,
		httpClient: client,
		orgId:      orgId,
		userAgent:  "terraform-provider-snyk",
	}
}

func (c *Client) doRequest(method, endpoint string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.ctx, method, fmt.Sprintf("%s/%s", c.hostUrl, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	return b, nil
}

func (c Client) Validate() error {
	return nil
}
