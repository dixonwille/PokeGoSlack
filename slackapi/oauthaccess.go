package slackapi

import (
	"encoding/json"
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
	w.Header().Add("Content-type", "text/html")
	io.Copy(w, res.Body)
	body := new(model.OAuthResp)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
