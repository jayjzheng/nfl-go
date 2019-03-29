package api

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Drives map[string]Drive

func (dd *Drives) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Wrap(err, "unmarshal raw")
	}

	m := make(Drives)
	var d Drive
	for k, v := range raw {
		if k == "crntdrv" {
			continue // TODO
		}

		if err := json.Unmarshal(v, &d); err != nil {
			return errors.Wrap(err, "unmarshal Drive")
		}

		m[k] = d
	}

	*dd = m
	return nil
}

type Drive struct {
	Number         string
	PossessionTeam string               `json:"posteam"`
	Quarter        int                  `json:"qtr"`
	Redzone        bool                 `json:"redzone"`
	Plays          map[json.Number]Play `json:"plays"`
	Fds            int                  `json:"fds"`
	Result         string               `json:"result"`
	PenaltyYards   int                  `json:"penyds"`
	YardsGained    int                  `json:"ydsgained"`
	NumberOfPlays  int                  `json:"numplays"`
	PossessionTime string               `json:"postime"`
	Start          driveInfo            `json:"start"`
	End            driveInfo            `json:"end"`
}

type driveInfo struct {
	Quarter  int    `json:"qtr"`
	Time     string `json:"time"`
	Yardline string `json:"yrdln"`
	Team     string `json:"team"`
}
