package api

import (
	"encoding/json"

	"github.com/jayjzheng/nfl-go"
)

type Play struct {
	ScoringPlay    int                     `json:"sp"`
	Quarter        int                     `json:"qtr"`
	Down           int                     `json:"down"`
	Time           string                  `json:"time"`
	Yardline       string                  `json:"yrdln"`
	YardsToGo      int                     `json:"ydstogo"`
	YardsNet       int                     `json:"ydsnet"`
	PossessionTeam string                  `json:"posteam"`
	Description    string                  `json:"desc"`
	Note           *string                 `json:"note"`
	Players        map[nfl.GSIS][]PlayStat `json:"players"`
}

type PlayStat struct {
	Sequence int         `json:"sequence"`
	Clubcode string      `json:"clubcode"`
	Name     string      `json:"playerName"`
	StatID   int         `json:"statId"`
	Yards    json.Number `json:"yards"`
}
