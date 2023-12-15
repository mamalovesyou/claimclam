// client.go
package podcastclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mamalovesyou/getclaim/graphql/gen/model"
)

const BaseURL = "https://601f1754b5a0e9001706a292.mockapi.io"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 30, // Add default 30 seconds timeout
		},
	}
}

func (c *Client) GetPodcasts(search, title, categoryName string, page, limit int) ([]model.Podcast, error) {
	endpoint := fmt.Sprintf("%s/podcasts", BaseURL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if search != "" {
		q.Add("search", search)
	}
	if title != "" {
		q.Add("title", title)
	}
	if categoryName != "" {
		q.Add("categoryName", categoryName)
	}
	if page > 0 {
		q.Add("page", fmt.Sprintf("%d", page))
	}
	if limit > 0 {
		q.Add("limit", fmt.Sprintf("%d", limit))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("request was canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	var podcasts []model.Podcast
	if err := json.NewDecoder(resp.Body).Decode(&podcasts); err != nil {
		return nil, err
	}

	return podcasts, nil
}
