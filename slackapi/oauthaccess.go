package slackapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/model"
)

const (
	oauthAccessEnd string = "/oauth.access"
)

//OAuthAccess is to ask for permission to connecet to slack
func OAuthAccess(w http.ResponseWriter, code string) (*model.OAuthResp, error) {
	params := url.Values{}
	params.Add("client_id", env.ClientID)
	params.Add("client_secret", env.ClientSecret)
	params.Add("code", code)
	api, err := newAPI(oauthAccessEnd, params)
	if err != nil {
		return nil, err
	}
	res, err := http.Get(api.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.Header.Get("Content-type") == "text/html" {
			io.Copy(w, res.Body)
			return nil, errors.New("Glitch")
		}
		return nil, errors.New(res.Header.Get("Content-type"))
	}
	if res.Header.Get("Content-type") != "applicaton/json" {
		return nil, errors.New("Can only accept json request")
	}
	body := new(model.OAuthResp)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
