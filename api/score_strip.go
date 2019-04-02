package api

import (
	"fmt"
)

type ScoreStripURLInput struct {
	Season, Week int
	SeasonType   string
	LiveUpdate   bool
}

func ScoreStripURL(in ScoreStripURLInput) string {
	uu := defaultBaseURLs()
	if in.LiveUpdate {
		return uu.LiveUpdateScoreStrip.String()
	}

	u := uu.ScoreStrip
	vv := u.Query()

	vv.Set("season", fmt.Sprintf("%d", in.Season))
	vv.Set("week", fmt.Sprintf("%d", in.Week))
	vv.Set("seasonType", in.SeasonType)

	u.RawQuery = vv.Encode()
	return u.String()
}

type ScoreStrip struct {
	Scores struct {
		Year int    `xml:"y,attr"`
		Week int    `xml:"w,attr"`
		Type string `xml:"t,attr"`
		GD   int    `xml:"gd,attr"` // ?
		BPH  string `xml:"bph,attr"`

		Games []struct {
			ID              string `xml:"eid,attr"`
			GSIS            string `xml:"gsis,attr"`
			Day             string `xml:"d,attr"`
			Time            string `xml:"t,attr"`
			Quarter         string `xml:"q,attr"`
			QuarterTimeLeft string `xml:"k,attr"`
			Home            string `xml:"h,attr"`
			HomeNickName    string `xml:"hnn,attr"`
			HomeScore       int    `xml:"hs,attr"`
			Visitor         string `xml:"v,attr"`
			VisitorNickName string `xml:"vnn,attr"`
			VisitorScore    int    `xml:"vs,attr"`
			Redzone         int    `xml:"rz,attr"`
			Possesion       string `xml:"p,attr"`
			GA              string `xml:"ga,attr"` // ?
			GameType        string `xml:"gt,attr"`
		} `xml:"g"`
	} `xml:"gms"`
}
