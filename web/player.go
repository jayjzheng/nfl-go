package web

import (
	"io"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"

	"github.com/jayjzheng/nfl-go"
	"github.com/pkg/errors"
)

const (
	baseURL = "http://www.nfl.com"
)

type Player struct {
	Href       string
	Name       string
	Position   string
	Number     *int // sometimes players don't have a number
	Status     string
	Height     string
	Weight     int
	BirthDate  string
	Experience int
	College    string
}

type PlayerIDs struct {
	GSIS nfl.GSIS
	ESB  nfl.ESB
}

func PlayerURL(href string) string {
	u, _ := url.Parse(baseURL)
	u.Path = href

	return u.String()
}

// ParsePlayerIDs parses text from r to find GSIS and ESB.
// it matches regex patterns: `GSIS ID: (\d+\-\d+)` and `ESB ID: (\S+)`
func ParsePlayerIDs(r io.Reader) (*PlayerIDs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "ReadAll")
	}

	ids := PlayerIDs{
		GSIS: nfl.GSIS(parse(b, `GSIS ID: (\d+\-\d+)`)),
		ESB:  nfl.ESB(parse(b, `ESB ID: (\S+)`)),
	}

	return &ids, nil
}

func parse(b []byte, rs string) string {
	re := regexp.MustCompile(rs)

	if ss := re.FindSubmatch(b); len(ss) > 1 {
		return strings.TrimSpace(string(ss[1]))
	}
	return ""
}
