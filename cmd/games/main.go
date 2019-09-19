package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

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
	ctx := context.Background()

	rr, err := api.GameRequests(ids)
	if err != nil {
		log.Fatalln(err)
	}

	resp := c.Do(ctx, rr, client.ValidateStatusOK)

	var games []api.Game
	fn := func(r *client.Response) error {
		defer r.Body.Close()

		var g api.Game
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
			return err
		}
		games = append(games, g)
		return nil
	}

	if err := c.Handle(ctx, resp, fn); err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, games, pretty); err != nil {
		log.Fatalln(err)
	}
}
