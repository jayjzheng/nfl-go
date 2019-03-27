package web

import (
	"io/ioutil"
	"regexp"
	"strings"

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
	GSIS string
	ESB  string
}

func (c *Client) FetchPlayerIDs(p Player) (*PlayerIDs, error) {
	u := c.BaseURL
	u.Path = p.Href

	resp, err := c.get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ReadAll")
	}

	var ids PlayerIDs

	gsis := regexp.MustCompile(`GSIS ID: (\d+\-\d+)`)
	if ss := gsis.FindSubmatch(b); len(ss) > 1 {
		ids.GSIS = strings.TrimSpace(string(ss[1]))
	}

	esb := regexp.MustCompile(`ESB ID: (.+)`)
	if ss := esb.FindSubmatch(b); len(ss) > 1 {
		ids.ESB = strings.TrimSpace(string(ss[1]))
	}

	if ids.GSIS == "" || ids.ESB == "" {
		return nil, ErrNotFound
	}

	return &ids, nil
}
