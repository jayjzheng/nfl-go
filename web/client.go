package web

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client struct {
	BaseURL *url.URL
	Http    interface {
		Do(*http.Request) (*http.Response, error)
	}
}

func NewClient(opts ...func(*Client)) *Client {
	var c Client

	for _, opt := range opts {
		opt(&c)
	}

	if c.Http == nil {
		c.Http = &http.Client{}
	}

	if c.BaseURL == nil {
		u, _ := url.Parse("http://www.nfl.com")
		c.BaseURL = u
	}

	return &c
}

func WithHttp(hc *http.Client) func(*Client) {
	return func(c *Client) {
		c.Http = hc
	}
}

func (c *Client) get(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest")
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Do")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("invalid status: %d", resp.StatusCode)
	}

	return resp, nil
}
