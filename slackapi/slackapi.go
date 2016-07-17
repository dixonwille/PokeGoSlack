package slackapi

import (
	"net/url"

	"github.com/dixonwille/PokeGoSlack/exception"
)

const (
	slackapi string = "https://slack.com/api/"
)

//NewAPI creates an endpoint to call
func NewAPI(endpoint string, params url.Values) (*url.URL, error) {
	api, err := url.Parse(slackapi)
	if err != nil {
		return nil, exception.NewInternalError("Could not parse the Slack Api: " + err.Error())
	}
	api.Path += endpoint
	api.RawQuery = params.Encode()
	return api, nil
}
