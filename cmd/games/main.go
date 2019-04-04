package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/jayjzheng/http-go/client"
	"github.com/jayjzheng/nfl-go/api"
	"github.com/jayjzheng/nfl-go/cmd"
)

const (
	sep = ","
)

var (
	idStr  string
	pretty bool
	c      client.Multi
)

func init() {
	flag.StringVar(&idStr, "ids", "", "comma seprated game ids")
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
	ids, err := cmd.SplitInput(idStr, sep)
	if err != nil {
		log.Fatalln(err)
	}

	resp := c.Get(api.GameURLs(ids), client.ValidateStatusOK)
	gg, err := games(resp, len(ids))
	if err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, gg, pretty); err != nil {
		log.Fatalln(err)
	}
}

func games(resp <-chan client.Response, count int) ([]api.Game, error) {
	var (
		errs  *multierror.Error
		games []api.Game
	)

	for i := 0; i < count; i++ {
		r := <-resp

		if r.Err != nil {
			errs = multierror.Append(errs, r.Err)
			continue
		}
		defer r.Body.Close()

		var g api.Game
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
			errs = multierror.Append(errs, err)
			continue
		}
		games = append(games, g)
	}

	return games, errs.ErrorOrNil()
}
