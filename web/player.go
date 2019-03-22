package web

import (
	"io/ioutil"
	"regexp"

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

func (c *Client) FetchPlayerGSIS(p Player) (string, error) {
	u := c.BaseURL
	u.Path = p.Href

	resp, err := c.get(u.String())
	if err != nil {
		return "", errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "ReadAll")
	}

	reg := regexp.MustCompile(`GSIS ID: (.+)`)
	if ss := reg.FindSubmatch(b); len(ss) > 1 {
		return string(ss[1]), nil
	}

	return "", ErrNotFound
}
