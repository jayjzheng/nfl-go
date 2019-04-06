package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jayjzheng/nfl-go/web"

	"github.com/jayjzheng/http-go/client"
	"github.com/jayjzheng/nfl-go/cmd"
)

const (
	sep = ","
)

var (
	urls       string
	pretty     bool
	concurrent int
	c          client.Multi
)

func init() {
	flag.StringVar(&urls, "urls", "", "comma seprated player urls")
	flag.BoolVar(&pretty, "pretty", false, "pretty print")
	flag.IntVar(&concurrent, "concurrent", 0, "concurrent")
	flag.Parse()

	c = client.Multi{
		ConcurrencyLimit: concurrent,
		Client: &http.Client{
			// do not follow redirects
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func main() {
	uu, err := cmd.SplitInput(urls, sep)
	if err != nil {
		log.Fatalln(err)
	}
	var (
		mu  sync.Mutex
		ctx = context.Background()
		ids = make(map[string]web.PlayerIDs)
	)

	resp := c.Get(ctx, uu, client.ValidateStatusOK)

	fn := func(r *client.Response) error {
		defer r.Body.Close()

		id, err := web.DecodePlayerIDsHTML(r.Body)
		if err != nil {
			return err
		}

		mu.Lock()
		ids[r.Request.URL.String()] = *id
		mu.Unlock()

		return nil
	}

	if err := c.Handle(ctx, resp, fn); err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, ids, pretty); err != nil {
		log.Fatalln(err)
	}
}
