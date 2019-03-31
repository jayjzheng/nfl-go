package web

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type Roster struct {
	Team    string
	Players []Player
}

func (c *Client) FetchRoster(team string) (*Roster, error) {
	resp, err := c.get(c.rosterURL(team))
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "goquery.NewDocumentFromReader")
	}

	var (
		r = Roster{
			Team:    team,
			Players: []Player{},
		}
		errs *multierror.Error
	)

	doc.Find("table#result tbody tr").Each(func(_ int, s *goquery.Selection) {
		sel := s.Find("td")

		if validSelection(sel) {
			p, err := c.toPlayer(sel)
			if err != nil {
				errs = multierror.Append(errs, err)
				return
			}
			r.Players = append(r.Players, *p)
		}
	})

	return &r, errs.ErrorOrNil()
}

func validSelection(sel *goquery.Selection) bool {
	return sel.Length() == 9
}

func (c *Client) toPlayer(s *goquery.Selection) (*Player, error) {
	href, _ := s.Eq(1).Find("a").Attr("href")

	parse := func(s *goquery.Selection, i int) string {
		return strings.TrimSpace(s.Eq(i).Text())
	}

	num, err := atoi(parse(s, 0))
	if err != nil {
		return nil, err
	}
	weight, err := atoi(parse(s, 5))
	if err != nil {
		return nil, err
	}
	exp, err := atoi(parse(s, 7))
	if err != nil {
		return nil, err
	}

	return &Player{
		Href:       href,
		Number:     num,
		Name:       parse(s, 1),
		Position:   parse(s, 2),
		Status:     parse(s, 3),
		Height:     parse(s, 4),
		Weight:     *weight,
		BirthDate:  parse(s, 6),
		Experience: *exp,
		College:    parse(s, 8),
	}, nil
}

func (c *Client) rosterURL(team string) *url.URL {
	u := *c.BaseURL
	u.Path = "/teams/roster"

	vv := u.Query()
	vv.Set("team", team)
	u.RawQuery = vv.Encode()

	return &u
}

func atoi(s string) (*int, error) {
	if s == "" {
		return nil, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
