package web

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestFetchPlayerIDGetError(t *testing.T) {
	m := client.NewMock(
		http.StatusOK,
		nil,
		errors.New("some-error"),
	)

	c := Client{
		BaseURL: &url.URL{},
		Http:    m,
	}

	_, err := c.FetchPlayerIDs(Player{})
	assert.Error(t, err)
}

func TestFetchPlayerIDsNotFound(t *testing.T) {
	m := client.NewMock(
		http.StatusOK,
		ioutil.NopCloser(strings.NewReader("")),
		nil,
	)

	c := Client{
		BaseURL: &url.URL{},
		Http:    m,
	}

	_, err := c.FetchPlayerIDs(Player{})
	assert.Equal(t, ErrNotFound, err)
}

func TestFetchPlayerIDsOK(t *testing.T) {
	f, err := os.Open(fmt.Sprintf("./fixtures/ahkellowitherspoon.html"))
	if err != nil {
		t.Fatal(err)
	}

	m := client.NewMock(
		http.StatusOK,
		ioutil.NopCloser(f),
		nil,
	)

	c := Client{
		BaseURL: &url.URL{},
		Http:    m,
	}

	ids, err := c.FetchPlayerIDs(Player{})
	if assert.NoError(t, err) {
		assert.Equal(t, "00-0033783", ids.GSIS, "gsis")
		assert.Equal(t, "WIT145608", ids.ESB, "esb")
	}
}
