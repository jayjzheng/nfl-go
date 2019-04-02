package web

import (
	"encoding/json"
	"io"
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

// RosterURL return the url for a team's roster.
func RosterURL(team string) string {
	u, _ := url.Parse(baseURL)
	u.Path = "/teams/roster"

	vv := u.Query()
	vv.Set("team", team)
	u.RawQuery = vv.Encode()

	return u.String()
}

// DecodeRosterJSON decodes JSON from r into a Roster.
func DecodeRosterJSON(r io.Reader) (*Roster, error) {
	var roster Roster

	if err := json.NewDecoder(r).Decode(&roster); err != nil {
		return nil, errors.Wrap(err, "Decode")
	}

	return &roster, nil
}

// DecodeRosterHTML decodes HTML from r into a Roster.
func DecodeRosterHTML(r io.Reader) (*Roster, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, errors.Wrap(err, "NewDocumentFromReader")
	}

	var (
		roster Roster
		errs   *multierror.Error
	)

	doc.Find("meta#teamName").Each(func(_ int, s *goquery.Selection) {
		roster.Team, _ = s.Attr("content")
	})

	doc.Find("table#result tbody tr").Each(func(_ int, s *goquery.Selection) {
		sel := s.Find("td")

		if validPlayer(sel) {
			p, err := toPlayer(sel)
			if err != nil {
				errs = multierror.Append(errs, err)
				return
			}
			roster.Players = append(roster.Players, *p)
		}
	})

	return &roster, errs.ErrorOrNil()
}

// RosterHTMLtoJSON decodes HTML from r into a Roster, converts it to JSON and writes to w.
// pretty adds indentation.
func RosterHTMLtoJSON(r io.Reader, w io.Writer, pretty bool) error {
	roster, err := DecodeRosterHTML(r)
	if err != nil {
		return errors.Wrap(err, "DecodeRosterHTML")
	}

	encoder := json.NewEncoder(w)
	if pretty {
		encoder.SetIndent("", "\t")
	}

	return errors.Wrap(encoder.Encode(roster), "Encode")
}

func validPlayer(sel *goquery.Selection) bool {
	return sel.Length() == 9
}

func toPlayer(s *goquery.Selection) (*Player, error) {
	href, _ := s.Eq(1).Find("a").Attr("href") // ok if empty

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
