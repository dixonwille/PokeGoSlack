package slackapi

import (
	"net/url"

	"github.com/dixonwille/PokeGoSlack/exception"
)

const (
	slackapi string = "https://slack.com/api/"
)

func newAPI(endpoint string, params url.Values) (*url.URL, error) {
	api, err := url.Parse(slackapi)
	if err != nil {
		return nil, exception.NewInternalErr(106, "Could not parse the Slack Api")
	}
	api.Path = endpoint
	api.RawQuery = params.Encode()
	return api, nil
}
