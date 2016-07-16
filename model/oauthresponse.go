package model

import "encoding/json"

//OAuthResp is the response from a handshake to Slack
type OAuthResp struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
}

//ParseOAuthResp parses the response from slack
func ParseOAuthResp(resp []byte) (*OAuthResp, error) {
	authResp := new(OAuthResp)
	err := json.Unmarshal(resp, authResp)
	if err != nil {
		return nil, err
	}
	return authResp, nil
}
