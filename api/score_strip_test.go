package api

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jayjzheng/http-go/client"
	"github.com/stretchr/testify/assert"
)

func TestFetchScoreStrip(t *testing.T) {
	f, err := os.Open("./fixtures/score_strip.xml")
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

	out, err := c.FetchScoreStrip(&FetchScoreStripInput{})
	if assert.NoError(t, err) {
		assert.Equal(t, "149", out.Data.BPH)
		assert.Equal(t, 2018, out.Data.Year)
		assert.Equal(t, 17, out.Data.Week)
		assert.Equal(t, "R", out.Data.Type)
		assert.Equal(t, 1, out.Data.GD)
		if assert.Equal(t, 16, len(out.Data.Games)) {
			g := out.Data.Games[0]

			assert.Equal(t, "2018123001", g.ID, "ID")
			assert.Equal(t, "57808", g.GSIS, "GSIS")
			assert.Equal(t, "Sun", g.Day, "Day")
			assert.Equal(t, "1:00", g.Time, "Time")
			assert.Equal(t, "F", g.Quarter, "Quarter")
			assert.Equal(t, "BUF", g.Home, "Home")
			assert.Equal(t, "bills", g.HomeNickName, "HomeNickName")
			assert.Equal(t, 42, g.HomeScore, "HomeScore")
			assert.Equal(t, "MIA", g.Visitor, "Visitor")
			assert.Equal(t, "dolphins", g.VisitorNickName, "VisitorNickName")
			assert.Equal(t, 17, g.VisitorScore, "VisitorScore")
			assert.Equal(t, 0, g.Redzone, "Redzone")
			assert.Equal(t, "", g.QuarterTimeLeft, "QuarterTimeLeft")
			assert.Equal(t, "", g.Possesion, "Possesion")
			assert.Equal(t, "", g.GA, "GA")
			assert.Equal(t, "REG", g.GameType, "GameType")
		}
	}
}
