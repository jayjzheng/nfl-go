package web

import (
	"os"
	"testing"

	"github.com/jayjzheng/nfl-go"

	"github.com/stretchr/testify/assert"
)

func TestDecodePlayerHTML(t *testing.T) {
	f, err := os.Open("./fixtures/ahkellowitherspoon.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	ids, err := DecodePlayerIDsHTML(f)
	if assert.NoError(t, err) {
		assert.Equal(t, nfl.GSIS("00-0033783"), ids.GSIS, "gsis")
		assert.Equal(t, nfl.ESB("WIT145608"), ids.ESB, "esb")
	}
}
