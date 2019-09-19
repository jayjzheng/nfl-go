package api

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Game struct {
	ID         string
	NextUpdate int64

	Home           Team
	Away           Team
	Drives         Drives
	ScoreSummaries map[json.Number]ScoreSummary

	Weather        string
	Media          string
	Quarter        string
	Note           string
	Down           int
	YardsToGo      int
	Redzone        bool
	Clock          string
	PossessionTeam string
	Stadium        string
}

func (g *Game) UnmarshalJSON(b []byte) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(b, &raw); err != nil {
		return errors.Wrap(err, "unmarshal raw")
	}

	for k, v := range raw {
		if k == "nextupdate" {
			if err := json.Unmarshal(v, &g.NextUpdate); err != nil {
				return errors.Wrap(err, "unmarshal nextupdate")
			}
			continue
		}

		g.ID = k

		var gj gameJSON
		if err := json.Unmarshal(v, &gj); err != nil {
			return errors.Wrap(err, "unmarshal gameJSON")
		}

		str := func(s *string) string {
			if s == nil {
				return ""
			}
			return *s
		}
		g.Weather = str(gj.Weather)
		g.Media = str(gj.Media)
		g.Quarter = gj.Quarter
		g.Note = str(gj.Note)
		g.Down = gj.Down
		g.YardsToGo = gj.YardsToGo
		g.Redzone = gj.Redzone
		g.Clock = gj.Clock
		g.PossessionTeam = gj.PossessionTeam
		g.Stadium = str(gj.Stadium)
		g.ScoreSummaries = gj.ScoreSummaries
		g.Drives = gj.Drives
		g.Home = gj.Home
		g.Away = gj.Away
	}

	return nil
}

func (g Game) MarshalJSON() ([]byte, error) {
	str := func(s string) *string {
		if s == "" {
			return nil
		}
		return &s
	}

	gj := gameJSON{
		Weather:        str(g.Weather),
		Media:          str(g.Media),
		Quarter:        g.Quarter,
		Note:           str(g.Note),
		Down:           g.Down,
		YardsToGo:      g.YardsToGo,
		Redzone:        g.Redzone,
		Clock:          g.Clock,
		PossessionTeam: g.PossessionTeam,
		Stadium:        str(g.Stadium),
		ScoreSummaries: g.ScoreSummaries,
		Drives:         g.Drives,
		Home:           g.Home,
		Away:           g.Away,
	}

	raw := map[string]interface{}{
		g.ID:         gj,
		"nextupdate": g.NextUpdate,
	}

	return json.Marshal(raw)
}

type gameJSON struct {
	Home           Team                         `json:"home"`
	Away           Team                         `json:"away"`
	Drives         Drives                       `json:"drives"`
	ScoreSummaries map[json.Number]ScoreSummary `json:"scrsummary"`

	Weather        *string `json:"weather"`
	Media          *string `json:"media"`
	YL             string  `json:"yl"`
	Quarter        string  `json:"qtr"`
	Note           *string `json:"note"`
	Down           int     `json:"down"`
	YardsToGo      int     `json:"togo"`
	Redzone        bool    `json:"redzone"`
	Clock          string  `json:"clock"`
	PossessionTeam string  `json:"posteam"`
	Stadium        *string `json:"stadium"`
}
