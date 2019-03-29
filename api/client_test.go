package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	assert.Equal(t, defaultBaseURLs(), c.BaseURLs)
	assert.NotNil(t, c.Http)
}

func TestWithHttp(t *testing.T) {
	hc := http.DefaultClient

	c := NewClient(
		WithHttp(hc),
	)

	assert.Equal(t, hc, c.Http)
}

func TestWithBaseURLs(t *testing.T) {
	uu := BaseURLs{
		ScoreStrip: &url.URL{Path: "some-path"},
	}

	c := NewClient(
		WithBaseURLs(&uu),
	)

	assert.Equal(t, &uu, c.BaseURLs)
}

func TestClientGetError(t *testing.T) {
	m := client.NewMock(
		http.StatusOK,
		nil,
		errors.New("some-error"),
	)
	c := Client{Http: m}

	_, err := c.get(&url.URL{})
	assert.NotNil(t, err)
}

func TestClientInvalidStatus(t *testing.T) {
	m := client.NewMock(http.StatusNotFound, nil, nil)
	c := Client{Http: m}

	_, err := c.get(&url.URL{})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "404")
	}
}

func TestClientGetOK(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader("foo"))
	m := client.NewMock(
		http.StatusOK,
		body,
		nil,
	)
	c := Client{Http: m}

	resp, err := c.get(&url.URL{})
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, body, resp.Body)
	}
}
