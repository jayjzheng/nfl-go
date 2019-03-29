package api

import (
	"github.com/jayjzheng/nfl-go"
)

type Stats struct {
	Passing    map[nfl.GSIS]PassingStat   `json:"passing"`
	Rushing    map[nfl.GSIS]RushingStat   `json:"rushing"`
	Receiving  map[nfl.GSIS]ReceivingStat `json:"receiving"`
	Fumbles    map[nfl.GSIS]FumbleStat    `json:"fumbles"`
	Kicking    map[nfl.GSIS]KickingStat   `json:"kicking"`
	Punting    map[nfl.GSIS]PuntingStat   `json:"punting"`
	KickReturn map[nfl.GSIS]ReturnStat    `json:"kickret"`
	PuntReturn map[nfl.GSIS]ReturnStat    `json:"puntret"`
	Defense    map[nfl.GSIS]DefenseStat   `json:"defense"`
	Team       TeamStat                   `json:"team"`
}

type PassingStat struct {
	PlayerName string `json:"name"`

	Attempt       int `json:"att"`
	Completions   int `json:"cmp"`
	Yards         int `json:"yds"`
	Touchdowns    int `json:"tds"`
	Interceptions int `json:"ints"`

	TwoPointAttempt int `json:"twopta"`
	TwoPointMade    int `json:"twoptm"`
}

type RushingStat struct {
	PlayerName string `json:"name"`

	Attempt    int `json:"att"`
	Yards      int `json:"yds"`
	Touchdowns int `json:"tds"`
	Long       int `json:"lng"`

	TwoPointAttempt int `json:"twopta"`
	TwoPointMade    int `json:"twoptm"`
}

type ReceivingStat struct {
	PlayerName string `json:"name"`

	Receptions int `json:"rec"`
	Yards      int `json:"yds"`
	Touchdowns int `json:"tds"`
	Long       int `json:"lng"`

	TwoPointAttempt int `json:"twopta"`
	TwoPointMade    int `json:"twoptm"`
}

type FumbleStat struct {
	PlayerName string `json:"name"`

	Total          int `json:"tot"`
	Recovered      int `json:"rcv"`
	TotalRecovered int `json:"trcv"`
	Yards          int `json:"yds"`
	Lost           int `json:"lost"`
}

type KickingStat struct {
	PlayerName string `json:"name"`

	FieldGoalsMade      int `json:"fgm"`
	FieldGoalsAttempted int `json:"fga"`
	FieldGoalYards      int `json:"fgyds"`
	FieldGoalTotal      int `json:"totpfg"`

	ExtraPointMade      int `json:"xpmade"`
	ExtraPointMisses    int `json:"xpmissed"`
	ExtraPointAttempted int `json:"xpa"`
	ExtraPointBlocked   int `json:"xpb"`
	ExtraPointTotal     int `json:"xptot"`
}

type PuntingStat struct {
	PlayerName string `json:"name"`

	Punts    int `json:"pts"`
	Yards    int `json:"yds"`
	Average  int `json:"avg"`
	Inside20 int `json:"i20"`
	Long     int `json:"lng"`
}

type ReturnStat struct {
	PlayerName string `json:"name"`

	Returns    int `json:"ret"`
	Average    int `json:"avg"`
	Touchdowns int `json:"tds"`
	Long       int `json:"lng"`
}

type DefenseStat struct {
	PlayerName string `json:"name"`

	Tackles       int     `json:"tkl"`
	Assists       int     `json:"ast"`
	Sacks         float64 `json:"sk"`
	Interceptions int     `json:"int"`
	ForcedFumbles int     `json:"ffum"`
}

type TeamStat struct {
	TotalFirstDowns  int    `json:"totfd"`
	TotalYards       int    `json:"totyds"`
	PassingYards     int    `json:"pyds"`
	RushingYards     int    `json:"ryds"`
	Penalties        int    `json:"pen"`
	PenaltyYards     int    `json:"penyds"`
	Turnovers        int    `json:"trnovr"`
	Punts            int    `json:"pt"`
	PuntYards        int    `json:"ptyds"`
	PuntAverage      int    `json:"ptavg"`
	TimeOfPossession string `json:"top"`
}
