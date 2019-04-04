package web

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRosterURLs(t *testing.T) {
	uu := RosterURLs([]string{
		"foo",
		"bar",
	})

	exp := []string{
		"http://www.nfl.com/teams/roster?team=foo",
		"http://www.nfl.com/teams/roster?team=bar",
	}

	assert.Equal(t, exp, uu)
}

func TestRosterURL(t *testing.T) {
	u := RosterURL("foo")

	assert.Equal(t, "http://www.nfl.com/teams/roster?team=foo", u)
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
			Href:       "/player/arikarmstead/2552493/profile",
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
