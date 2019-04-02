package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jayjzheng/http-go/client"

	"github.com/jayjzheng/nfl-go"
	"github.com/jayjzheng/nfl-go/cmd"
	"github.com/jayjzheng/nfl-go/web"
)

const (
	sep = ","
)

var (
	teams  string
	pretty bool
	c      client.Multi
)

func init() {
	flag.StringVar(&teams, "teams", "", "comma seprated team abbreviations")
	flag.BoolVar(&pretty, "pretty", false, "pretty print")
	flag.Parse()

	if cmd.PipedInput() {
		teams += sep + string(cmd.ScanStdin())
	}

	c = client.Multi{
		Client: &http.Client{
			// do not follow redirects
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func main() {
	tt := cmd.StringSlice(teams, sep)
	if len(tt) == 0 {
		tt = nfl.TeamAbbreviations
	}

	urls := make([]string, len(tt))
	for i, t := range tt {
		urls[i] = web.RosterURL(t)
	}

	rr, err := c.Get(urls, client.ValidateStatusOK)
	if err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, rosters(rr), pretty); err != nil {
		log.Fatalln(err)
	}
}

func rosters(rr []*http.Response) []web.Roster {
	var rosters []web.Roster
	for _, r := range rr {
		defer r.Body.Close()

		roster, err := web.DecodeRosterHTML(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		rosters = append(rosters, *roster)
	}
	return rosters
}
