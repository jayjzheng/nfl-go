package main

import (
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jayjzheng/nfl-go/api"
	"github.com/jayjzheng/nfl-go/cmd"
)

var (
	season, week int
	seasonType   string
	live, pretty bool
	c            *http.Client
)

func init() {
	flag.IntVar(&season, "season", 2018, "season")
	flag.IntVar(&week, "week", 1, "week")
	flag.StringVar(&seasonType, "type", "REG", "season type")
	flag.BoolVar(&pretty, "pretty", false, "pretty print")
	flag.BoolVar(&live, "live", false, "live update")
	flag.Parse()

	c = &http.Client{}
}

func main() {
	u := api.ScoreStripURL(api.ScoreStripURLInput{
		Season:     season,
		Week:       week,
		SeasonType: seasonType,
		LiveUpdate: live,
	})

	resp, err := c.Get(u)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var ss api.ScoreStrip
	if err := xml.NewDecoder(resp.Body).Decode(&ss); err != nil {
		log.Fatalln(err)
	}

	if err := cmd.WriteJSON(os.Stdout, ss.Scores, pretty); err != nil {
		log.Fatalln(err)
	}
}
