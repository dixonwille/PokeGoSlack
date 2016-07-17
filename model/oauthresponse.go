package model

import (
	"encoding/json"

	"github.com/dixonwille/PokeGoSlack/exception"
)

//OAuthResp is the response from a handshake to Slack
type OAuthResp struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
	Error       string `json:"error"`
}

//ParseOAuthResp parses the response from slack
func ParseOAuthResp(resp []byte) (*OAuthResp, error) {
	authResp := new(OAuthResp)
	err := json.Unmarshal(resp, authResp)
	if err != nil {
		return nil, exception.NewInternalError(err.Error())
	}
	return authResp, nil
}
