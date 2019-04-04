package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/go-multierror"

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

	resp := c.Get(web.RosterURLs(teams), client.ValidateStatusOK)
	rr, err := rosters(resp, len(teams))
	if err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, rr, pretty); err != nil {
		log.Fatalln(err)
	}
}

func rosters(resp <-chan client.Response, count int) ([]web.Roster, error) {
	var (
		errs    *multierror.Error
		rosters []web.Roster
	)

	for i := 0; i < count; i++ {
		r := <-resp

		if r.Err != nil {
			errs = multierror.Append(errs, r.Err)
			continue
		}
		defer r.Body.Close()

		roster, err := web.DecodeRosterHTML(r.Body)
		if err != nil {
			errs = multierror.Append(errs, r.Err)
			continue
		}
		rosters = append(rosters, *roster)
	}

	return rosters, errs.ErrorOrNil()
}