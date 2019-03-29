package web

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/jayjzheng/nfl-go"
	"github.com/pkg/errors"
)

type Player struct {
	Href       string
	Name       string
	Position   string
	Number     string
	Status     string
	Height     string
	Weight     string
	BirthDate  string
	Experience string
	College    string
}

type PlayerIDs struct {
	GSIS nfl.GSIS
	ESB  nfl.ESB
}

func (c *Client) FetchPlayerIDs(p Player) (*PlayerIDs, error) {
	u := c.BaseURL
	u.Path = p.Href

	resp, err := c.get(u)
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ReadAll")
	}

	ids := PlayerIDs{
		GSIS: nfl.GSIS(parse(b, `GSIS ID: (\d+\-\d+)`)),
		ESB:  nfl.ESB(parse(b, `ESB ID: (.+)`)),
	}

	if ids.GSIS == "" || ids.ESB == "" {
		return nil, ErrNotFound
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
