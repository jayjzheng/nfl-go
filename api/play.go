package api

import (
	"github.com/jayjzheng/nfl-go"
)

type Play struct {
	ScoringPlay    int    `json:"sp"`
	Quarter        int    `json:"qtr"`
	Down           int    `json:"down"`
	Time           string `json:"time"`
	Yardline       string `json:"yrdln"`
	YardsToGo      int    `json:"ydstogo"`
	YardsNet       int    `json:"ydsnet"`
	PossessionTeam string `json:"posteam"`
	Description    string `json:"desc"`
	Note           string `json:"note"`
	Players        map[nfl.GSIS][]struct {
		Sequence int    `json:"sequence"`
		Clubcode string `json:"clubcode"`
		Name     string `json:"playerName"`
		StatID   int    `json:"statId"`
		Yards    int    `json:"yards"`
	} `json:"players"`
}
