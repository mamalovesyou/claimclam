package podcasts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/mamalovesyou/getclaim/graphql/gen/model"
	"github.com/mamalovesyou/getclaim/internal/logging"
	"go.uber.org/zap"
)

const BaseURL = "https://601f1754b5a0e9001706a292.mockapi.io"

var (
	ErrInvalidPageParams  = errors.New("invalid parameter: page must be greater than 0")
	ErrInvalidLimitParams = errors.New("invalid parameter: limit must be greater than 0")
)

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

type ListPodcastsParams struct {
	Search       *string
	Title        *string
	CategoryName *string
	Page         *int
	Limit        *int
}

func (p *ListPodcastsParams) Validate() error {
	// If page number is not provided or is invalid, set it to 1
	if p.Page != nil && pointer.GetInt(p.Page) < 1 {
		return ErrInvalidPageParams
	}
	// If limit is not provided or is invalid, set it to 10
	if p.Limit != nil && pointer.GetInt(p.Limit) < 1 {
		return ErrInvalidLimitParams
	}
	return nil
}

func (p *ListPodcastsParams) UpdateQueryParams(q url.Values) url.Values {
	if p.Search != nil {
		q.Add("search", pointer.GetString(p.Search))
	}
	if p.Title != nil {
		q.Add("title", pointer.GetString(p.Title))
	}
	if p.CategoryName != nil {
		q.Add("categoryName", pointer.GetString(p.CategoryName))
	}
	if p.Page != nil {
		q.Add("page", fmt.Sprintf("%d", pointer.GetInt(p.Page)))
	}
	if p.Limit != nil {
		q.Add("limit", fmt.Sprintf("%d", pointer.GetInt(p.Limit)))
	}
	return q
}

func (c *Client) ListPodcasts(ctx context.Context, params *ListPodcastsParams) ([]*model.Podcast, error) {

	// Verify Params for query
	if err := params.Validate(); err != nil {
		logging.WithContext(ctx).Error("Invalid params", zap.Error(err))
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/podcasts", BaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q = params.UpdateQueryParams(q)
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

	var podcasts []*model.Podcast
	if err := json.NewDecoder(resp.Body).Decode(&podcasts); err != nil {
		return nil, err
	}

	return podcasts, nil
}
