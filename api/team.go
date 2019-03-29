package api

import "encoding/json"

type Team struct {
	Score struct {
		FirstQuarter  int `json:"1"`
		SecondQuarter int `json:"2"`
		ThirdQuarter  int `json:"3"`
		FourthQuarter int `json:"4"`
		Overtime      int `json:"5"`
		Total         int `json:"T"`
	} `json:"score"`

	Abbreviation string          `json:"abbr"`
	TurnOvers    int             `json:"to"`
	Players      json.RawMessage `json:"players"` // ?
	Stats        Stats           `json:"stats"`
}
