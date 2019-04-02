package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreStripURL(t *testing.T) {
	u := ScoreStripURL(ScoreStripURLInput{
		LiveUpdate: true,
	})
	assert.Equal(t, "http://www.nfl.com/liveupdate/scorestrip/ss.xml", u)

	u = ScoreStripURL(ScoreStripURLInput{
		Season:     2009,
		Week:       1,
		SeasonType: "foo",
	})
	assert.Equal(t, "http://www.nfl.com/ajax/scorestrip?season=2009&seasonType=foo&week=1", u)
}
