package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreStripRequest(t *testing.T) {
	req, err := ScoreStripRequest(ScoreStripRequestInput{
		LiveUpdate: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, "http://www.nfl.com/liveupdate/scorestrip/ss.xml", req.URL.String())
	assert.Equal(t, http.MethodGet, req.Method)

	req, err = ScoreStripRequest(ScoreStripRequestInput{
		Season:     2009,
		Week:       1,
		SeasonType: "foo",
	})
	assert.NoError(t, err)
	assert.Equal(t, "http://www.nfl.com/ajax/scorestrip?season=2009&seasonType=foo&week=1", req.URL.String())
	assert.Equal(t, http.MethodGet, req.Method)
}

func TestGameRequests(t *testing.T) {
	rr, err := GameRequests([]string{
		"foo",
		"bar",
	})
	assert.NoError(t, err)

	uu := []string{
		"http://www.nfl.com/liveupdate/game-center/foo/foo_gtd.json",
		"http://www.nfl.com/liveupdate/game-center/bar/bar_gtd.json",
	}

	for i, u := range uu {
		req := rr[i]
		assert.Equal(t, u, req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
	}
}

func TestGameRequest(t *testing.T) {
	req, err := GameRequest("foo")
	assert.NoError(t, err)

	assert.Equal(t,
		"http://www.nfl.com/liveupdate/game-center/foo/foo_gtd.json",
		req.URL.String(),
	)
	assert.Equal(t, http.MethodGet, req.Method)
}
