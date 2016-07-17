package slackapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

const (
	oauthAccessEnd string = "oauth.access"
)

//OAuthAccess calls Slack with code to verify access
func OAuthAccess(code string) (*model.OAuthResp, error) {
	params := url.Values{}
	params.Add("client_id", env.ClientID)
	params.Add("client_secret", env.ClientSecret)
	params.Add("code", code)
	api, err := NewAPI(oauthAccessEnd, params)
	if err != nil {
		return nil, err
	}
	res, err := http.Get(api.String())
	if err != nil {
		return nil, exception.NewInternalError(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, exception.NewInternalError("Something went wrong with slack")
	}
	if !strings.Contains(res.Header.Get("Content-type"), "application/json") {
		return nil, exception.NewInternalError("Slack sent a response back but it should be application/json")
	}
	body := new(model.OAuthResp)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		return nil, exception.NewInternalError(err.Error())
	}
	return body, nil
}
