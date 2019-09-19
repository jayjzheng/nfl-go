package main

import (
	"context"
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
	teamStr string
	pretty  bool
	c       client.Multi
)

func init() {
	flag.StringVar(&teamStr, "teams", "", "comma seprated team abbreviations")
	flag.BoolVar(&pretty, "pretty", false, "pretty print")
	flag.Parse()

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
	teams, err := cmd.SplitInput(teamStr, sep)
	if err != nil {
		log.Fatalln(err)
	}
	if len(teams) == 0 {
		teams = nfl.TeamAbbreviations
	}
	ctx := context.Background()

	rr, err := web.RosterRequests(teams)
	if err != nil {
		log.Fatalln(err)
	}

	resp := c.Do(ctx, rr, client.ValidateStatusOK)

	var rosters []web.Roster
	fn := func(r *client.Response) error {
		defer r.Body.Close()

		roster, err := web.DecodeRosterHTML(r.Body)
		if err != nil {
			return err
		}
		rosters = append(rosters, *roster)
		return nil
	}

	if err := c.Handle(ctx, resp, fn); err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, rosters, pretty); err != nil {
		log.Fatalln(err)
	}
}
