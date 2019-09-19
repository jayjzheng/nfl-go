package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameUnmarshalJSON(t *testing.T) {
	f, err := os.Open("./fixtures/2012080953.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	var g Game
	err = g.UnmarshalJSON(b)

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
		assert.Equal(t, 26, g.Drives.Current, "current drives")
		assert.Equal(t, 26, len(g.Drives.Drives), "number of drives")
		assert.Equal(t, 3, len(g.ScoreSummaries), "number of score summaries")
	}
}

func TestGameMarshalJSON(t *testing.T) {
	f1, err := os.Open("./fixtures/2012080953.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()

	exp, err := ioutil.ReadAll(f1)
	if err != nil {
		t.Fatal(err)
	}

	var g Game
	if err := json.Unmarshal(exp, &g); err != nil {
		t.Fatal(err)
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(g); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(exp), b.Len())
}
