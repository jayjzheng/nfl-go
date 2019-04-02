package web

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRosterURL(t *testing.T) {
	u := RosterURL("foo")

	assert.Equal(t, "http://www.nfl.com/teams/roster?team=foo", u)
}

func TestDecodeRosterJSON(t *testing.T) {
	f, err := os.Open("./fixtures/sf.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r, err := DecodeRosterJSON(f)
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

func TestRosterHTMLtoJSON(t *testing.T) {
	f, err := os.Open("./fixtures/sf.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var b1 bytes.Buffer
	err = RosterHTMLtoJSON(f, &b1, false)

	assert.NoError(t, err)
	assert.Equal(t, 14361, b1.Len())

	if _, err := f.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	var b2 bytes.Buffer
	err = RosterHTMLtoJSON(f, &b2, true)

	assert.NoError(t, err)
	assert.Equal(t, 18234, b2.Len())
}
