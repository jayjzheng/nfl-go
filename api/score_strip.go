package api

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type FetchScoreStripInput struct {
	Week       int
	Season     int
	SeasonType string
	LiveUpdate bool
}

type FetchScoreStripOutput struct {
	Data struct {
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

func (c *Client) scoreStripURL(in *FetchScoreStripInput) *url.URL {
	if in.LiveUpdate {
		return c.BaseURLs.LiveUpdateScoreStrip
	}

	u := c.BaseURLs.ScoreStrip
	vv := u.Query()

	vv.Set("season", fmt.Sprintf("%d", in.Season))
	vv.Set("week", fmt.Sprintf("%d", in.Week))
	vv.Set("seasonType", in.SeasonType)

	u.RawQuery = vv.Encode()
	return u
}

func (c *Client) FetchScoreStrip(in *FetchScoreStripInput) (*FetchScoreStripOutput, error) {
	u := c.scoreStripURL(in)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest")
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Http.Do")
	}
	defer resp.Body.Close()

	var out FetchScoreStripOutput
	if err := xml.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	return &out, nil
}