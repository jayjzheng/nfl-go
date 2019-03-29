package api

import "github.com/jayjzheng/nfl-go"

type ScoreSummary struct {
	Type        string              `json:"type"`
	Description string              `json:"desc"`
	Quarter     int                 `json:"qtr"`
	Team        string              `json:"team"`
	Players     map[string]nfl.GSIS `json:"players"`
}
