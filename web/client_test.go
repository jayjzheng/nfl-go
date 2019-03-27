package web

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	assert.Equal(t, "http://www.nfl.com", c.BaseURL.String())
	assert.NotNil(t, c.Http)
}

func TestWithHttp(t *testing.T) {
	hc := http.DefaultClient

	c := NewClient(
		WithHttp(hc),
	)

	assert.Equal(t, hc, c.Http)
}

func TestClientGetError(t *testing.T) {
	m := client.NewMock(
		http.StatusOK,
		nil,
		errors.New("some-error"),
	)
	c := Client{Http: m}

	_, err := c.get("some-url")
	assert.NotNil(t, err)
}

func TestClientInvalidStatus(t *testing.T) {
	m := client.NewMock(http.StatusNotFound, nil, nil)
	c := Client{Http: m}

	_, err := c.get("some-url")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "404")
	}
}
