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

func TestFetchPlayers(t *testing.T) {
	c := Client{
		BaseURL: &url.URL{},
		Http:    mockRoster(t, "sf"),
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

func mockRoster(t *testing.T, team string) *client.Mock {
	f, err := os.Open(fmt.Sprintf("./fixtures/rosters/%s.html", team))
	if err != nil {
		t.Fatal(err)
	}

	return client.NewMock(
		http.StatusOK,
		ioutil.NopCloser(f),
		nil,
	)
}
