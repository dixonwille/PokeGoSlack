package slackapi

import (
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"
	"github.com/dixonwille/PokeGoSlack/env"
)

const (
	oauthAccessEnd string = "/oauth.access"
)

//OAuthAccess is to ask for permission to connecet to slack
func OAuthAccess(code string) error {
	params := url.Values{}
	params.Add("client_id", env.ClientID)
	params.Add("client_secret", env.ClientSecret)
	params.Add("code", code)
	api, err := newAPI(oauthAccessEnd, params)
	if err != nil {
		return err
	}
	res, err := http.Get(api.String())
	if err != nil {
		return err
	}
	spew.Dump(res.Body)
	return nil
}
