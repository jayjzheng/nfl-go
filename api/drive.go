package api

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Drives struct {
	Current int
	Drives  map[json.Number]Drive
}

func (dd *Drives) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Wrap(err, "unmarshal raw")
	}

	dd.Drives = make(map[json.Number]Drive)
	for k, v := range raw {
		if k == "crntdrv" {
			if err := json.Unmarshal(v, &dd.Current); err != nil {
				return errors.Wrap(err, "unmarshal nextupdate")
			}
			continue
		}

		var d Drive
		if err := json.Unmarshal(v, &d); err != nil {
			return errors.Wrap(err, "unmarshal Drive")
		}
		dd.Drives[json.Number(k)] = d
	}

	return nil
}

func (dd Drives) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{
		"crntdrv": dd.Current,
	}

	for k, d := range dd.Drives {
		raw[string(k)] = d
	}

	return json.Marshal(raw)
}

type Drive struct {
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
