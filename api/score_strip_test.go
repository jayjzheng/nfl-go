package api

import (
	"fmt"
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
		fmt.Println(out)
	}
}
