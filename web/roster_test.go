package web

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestFetchRosterGetError(t *testing.T) {

	m := client.NewMock(
		http.StatusOK,
		nil,
		errors.New("some-error"),
	)

	c := Client{
		BaseURL: &url.URL{},
		Http:    m,
	}

	_, err := c.FetchRoster("")
	assert.Error(t, err)
}

func TestFetchRosterOK(t *testing.T) {
	f, err := os.Open("./fixtures/sf.html")
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

	r, err := c.FetchRoster("SF")
	if assert.NoError(t, err) {
		assert.Equal(t, "SF", r.Team)
		assert.Equal(t, len(r.Players), 69)
		assert.EqualValues(t, Player{
			Href:       "/player/kwonalexander/2552592/profile",
			Name:       "Alexander, Kwon",
			Position:   "MLB",
			Status:     "ACT",
			Height:     "6'1\"",
			Weight:     "227",
			BirthDate:  "8/3/1994",
			Experience: "5",
			College:    "LSU",
		}, r.Players[0])
	}
}
