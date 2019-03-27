package api

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client struct {
	BaseURLs *BaseURLs
	Http     interface {
		Do(*http.Request) (*http.Response, error)
	}
}

type BaseURLs struct {
	ScoreStrip           *url.URL
	LiveUpdateScoreStrip *url.URL
	LiveUpdateGame       *url.URL
}

func NewClient(opts ...func(*Client)) *Client {
	var c Client

	for _, opt := range opts {
		opt(&c)
	}

	if c.Http == nil {
		c.Http = &http.Client{}
	}

	if c.BaseURLs == nil {
		c.BaseURLs = defaultBaseURLs()
	}

	return &c
}

func WithHttp(hc *http.Client) func(*Client) {
	return func(c *Client) {
		c.Http = hc
	}
}

func WithBaseURLs(bu *BaseURLs) func(*Client) {
	return func(c *Client) {
		c.BaseURLs = bu
	}
}

func (c *Client) get(u string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
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

func defaultBaseURLs() *BaseURLs {
	return &BaseURLs{
		LiveUpdateScoreStrip: mustParseURL("http://www.nfl.com/liveupdate/scorestrip/ss.xml"),
		LiveUpdateGame:       mustParseURL("http://www.nfl.com/liveupdate/game-center"),
		ScoreStrip:           mustParseURL("http://www.nfl.com/ajax/scorestrip"),
	}
}

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
