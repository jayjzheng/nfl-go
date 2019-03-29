package api

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestFetchGame(t *testing.T) {
	f, err := os.Open("./fixtures/2012080953.json")
	if err != nil {
		t.Fatal(err)
	}

	m := client.NewMock(
		http.StatusOK,
		ioutil.NopCloser(f),
		nil,
	)

	c := Client{
		BaseURLs: defaultBaseURLs(),
		Http:     m,
	}

	g, err := c.FetchGame("2012080953")
	if assert.NoError(t, err) {
		assert.Equal(t, "2012080953", g.ID, "ID")
		assert.Equal(t, int64(289), g.NextUpdate, "NextUpdate")
		assert.Equal(t, "Final", g.Quarter, "Quarter")
		assert.Equal(t, 0, g.Down, "Down")
		assert.Equal(t, 0, g.YardsToGo, "YardsToGo")
		assert.Equal(t, true, g.Redzone, "Redzone")
		assert.Equal(t, "00:27", g.Clock, "Clock")
		assert.Equal(t, "NE", g.PossessionTeam, "PossessionTeam")
		assert.Equal(t, "", g.Weather, "Weather")
		assert.Equal(t, "", g.Media, "Media")
		assert.Equal(t, "", g.Note, "Note")
		assert.Equal(t, "", g.Stadium, "Stadium")
		assert.Equal(t, "NE", g.Home.Abbreviation)
		assert.Equal(t, "NO", g.Away.Abbreviation)
		assert.Equal(t, 26, len(g.Drives), "number of drives")
		assert.Equal(t, 3, len(g.ScoreSummaries), "number of score summaries")
	}
}

func TestGameURL(t *testing.T) {
	uu := defaultBaseURLs()
	c := Client{BaseURLs: uu}

	u := c.gameURL("2012080953")
	assert.Equal(t,
		"http://www.nfl.com/liveupdate/game-center/2012080953/2012080953_gtd.json",
		u.String(),
	)
	assert.Equal(t, *uu, *c.BaseURLs)
}
