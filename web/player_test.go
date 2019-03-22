package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestFetchPlayerGSIS(t *testing.T) {
	c := Client{
		BaseURL: &url.URL{},
		Http:    mockPlayer(t, "ahkellowitherspoon"),
	}

	gsis, err := c.FetchPlayerGSIS(Player{})
	if assert.NoError(t, err) {
		assert.Equal(t, "00-0033783", gsis)
	}
}

func mockPlayer(t *testing.T, player string) *client.Mock {
	f, err := os.Open(fmt.Sprintf("./fixtures/players/%s.html", player))
	if err != nil {
		t.Fatal(err)
	}

	return client.NewMock(
		http.StatusOK,
		ioutil.NopCloser(f),
		nil,
	)
}
