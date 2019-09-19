package web

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRosterRequests(t *testing.T) {
	rr, err := RosterRequests([]string{
		"foo",
		"bar",
	})
	assert.NoError(t, err)

	uu := []string{
		"http://www.nfl.com/teams/roster?team=foo",
		"http://www.nfl.com/teams/roster?team=bar",
	}

	for i, u := range uu {
		req := rr[i]
		assert.Equal(t, u, req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
	}

}

func TestDecodeRosterHTML(t *testing.T) {
	f, err := os.Open("./fixtures/sf.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r, err := DecodeRosterHTML(f)
	if assert.NoError(t, err) {
		num := 91

		assert.Equal(t, "SF", r.Team)
		assert.Equal(t, len(r.Players), 69)
		assert.EqualValues(t, Player{
			URL:        "http://www.nfl.com/player/arikarmstead/2552493/profile",
			Name:       "Armstead, Arik",
			Number:     &num,
			Position:   "DE",
			Status:     "ACT",
			Height:     "6'7\"",
			Weight:     292,
			BirthDate:  "11/15/1993",
			Experience: 5,
			College:    "Oregon",
		}, r.Players[1])
	}
}
