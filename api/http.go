package api

import (
	"fmt"
	"net/http"
	"net/url"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type ScoreStripRequestInput struct {
	Season, Week int
	SeasonType   string
	LiveUpdate   bool
}

// ScoreStripRequest returns the http request for fetching a score strip.
func ScoreStripRequest(in ScoreStripRequestInput) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, scoreStripURL(in), nil)
}

// GameRequests returns a slice of http requests based on the ids.
func GameRequests(ids []string) ([]*http.Request, error) {
	var (
		rr   []*http.Request
		errs *multierror.Error
	)

	for _, id := range ids {
		req, err := GameRequest(id)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, id))
			continue
		}
		rr = append(rr, req)
	}

	return rr, errs.ErrorOrNil()
}

// GameRequest returns the http request for fetching a game based on the id.
func GameRequest(id string) (*http.Request, error) {
	u := fmt.Sprintf(
		"http://www.nfl.com/liveupdate/game-center/%s/%s_gtd.json",
		id, id,
	)
	return http.NewRequest(http.MethodGet, u, nil)
}

func scoreStripURL(in ScoreStripRequestInput) string {
	if in.LiveUpdate {
		return "http://www.nfl.com/liveupdate/scorestrip/ss.xml"
	}

	u, err := url.Parse("http://www.nfl.com/ajax/scorestrip")
	if err != nil {
		panic(err) // not expected
	}

	vv := u.Query()

	vv.Set("season", fmt.Sprintf("%d", in.Season))
	vv.Set("week", fmt.Sprintf("%d", in.Week))
	vv.Set("seasonType", in.SeasonType)

	u.RawQuery = vv.Encode()
	return u.String()
}
