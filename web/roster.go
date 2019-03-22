package web

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

var Teams = []string{
	"ATL", "BAL", "BUF", "CAR", "CHI", "CIN", "CLE",
	"DAL", "DEN", "DET", "GB", "TEN", "HOU", "IND",
	"JAX", "KC", "LA", "OAK", "MIA", "MIN", "NE",
	"NO", "NYG", "NYJ", "PHI", "ARI", "PIT", "LAC",
	"SF", "SEA", "TB", "WAS",
}

type Roster struct {
	Team    string
	Players []Player
}

func (c *Client) FetchRoster(team string) (*Roster, error) {
	resp, err := c.get(c.rosterURL(team).String())
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "goquery.NewDocumentFromReader")
	}

	r := Roster{
		Team:    team,
		Players: []Player{},
	}

	doc.Find("table#result tbody tr").Each(func(_ int, s *goquery.Selection) {
		sel := s.Find("td")

		if validSelection(sel) {
			r.Players = append(r.Players, c.toPlayer(sel))
		}
	})

	return &r, nil
}

func validSelection(sel *goquery.Selection) bool {
	return sel.Length() == 9
}

func (c *Client) toPlayer(s *goquery.Selection) Player {
	href, _ := s.Eq(1).Find("a").Attr("href")

	return Player{
		Href:       href,
		Number:     strings.TrimSpace(s.Eq(0).Text()),
		Name:       strings.TrimSpace(s.Eq(1).Text()),
		Position:   strings.TrimSpace(s.Eq(2).Text()),
		Status:     strings.TrimSpace(s.Eq(3).Text()),
		Height:     strings.TrimSpace(s.Eq(4).Text()),
		Weight:     strings.TrimSpace(s.Eq(5).Text()),
		BirthDate:  strings.TrimSpace(s.Eq(6).Text()),
		Experience: strings.TrimSpace(s.Eq(7).Text()),
		College:    strings.TrimSpace(s.Eq(8).Text()),
	}
}

func (c *Client) rosterURL(team string) *url.URL {
	u := *c.BaseURL
	u.Path = "/teams/roster"

	vv := u.Query()
	vv.Set("team", team)
	u.RawQuery = vv.Encode()

	return &u
}

func (c *Client) playerURL(path string) *url.URL {
	u := *c.BaseURL
	u.Path = path
	return &u
}
