package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/pkg/errors"
)

// FetchGame fetches a Game based on game id.
func (c *Client) FetchGame(id string) (*Game, error) {
	resp, err := c.get(c.gameURL(id))
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	g, err := DecodeGame(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "decodeGame")
	}

	return g, nil
}

type Game struct {
	ID         string
	NextUpdate int64

	Home           Team                         `json:"home"`
	Away           Team                         `json:"away"`
	Drives         Drives                       `json:"drives"`
	ScoreSummaries map[json.Number]ScoreSummary `json:"scrsummary"`

	Weather        string `json:"weather"`
	Media          string `json:"media"`
	Quarter        string `json:"qtr"`
	Note           string `json:"note"`
	Down           int    `json:"down"`
	YardsToGo      int    `json:"togo"`
	Redzone        bool   `json:"redzone"`
	Clock          string `json:"clock"`
	PossessionTeam string `json:"posteam"`
	Stadium        string `json:"stadium"`
}

func DecodeGame(r io.Reader) (*Game, error) {
	var raw map[string]json.RawMessage

	if err := json.NewDecoder(r).Decode(&raw); err != nil {
		return nil, errors.Wrap(err, "Decode raw")
	}

	var (
		g    Game
		next int64
	)

	for k, v := range raw {
		if k == "nextupdate" {
			if err := json.Unmarshal(v, &next); err != nil {
				return nil, errors.Wrap(err, "Unmarshal nextupdate")
			}
			continue
		}
		if err := json.Unmarshal(v, &g); err != nil {
			return nil, errors.Wrap(err, "Unmarshal game")
		}
		g.ID = k
	}
	g.NextUpdate = next
	return &g, nil
}

func (c *Client) gameURL(id string) *url.URL {
	u := *c.BaseURLs.LiveUpdateGame
	u.Path += fmt.Sprintf("/%s/%s_gtd.json", id, id)
	return &u
}
