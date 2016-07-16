package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/dixonwille/PokeGoSlack/slackapi"
	"github.com/gorilla/context"
)

const (
	oauthAccessEnd string = "oauth.access"
)

//OAuthAccess is to ask for permission to connecet to slack
func OAuthAccess(w http.ResponseWriter, r *http.Request) {
	code, ok := context.Get(r, env.KeyCode).(string)
	if !ok {
		newErr := exception.NewInternalErr(105, "Cannot get code from context")
		errMsg := model.NewErrorMessage(newErr.Error())
		newErr.LogError()
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	params := url.Values{}
	params.Add("client_id", env.ClientID)
	params.Add("client_secret", env.ClientSecret)
	params.Add("code", code)
	api, err := slackapi.NewAPI(oauthAccessEnd, params)
	if err != nil {
		errMsg := model.NewErrorMessage(err.Error())
		env.Logger.Println(err.Error())
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	spew.Dump(api.String())
	res, err := http.Get(api.String())
	if err != nil {
		errMsg := model.NewErrorMessage(err.Error())
		env.Logger.Println(err.Error())
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	defer res.Body.Close()
	if strings.Contains(res.Header.Get("Content-type"), "text/html") {
		io.Copy(w, res.Body)
		return
	}
	if res.StatusCode != http.StatusOK {
		newErr := exception.NewInternalErr(107, "Something went wrong with slack")
		errMsg := model.NewErrorMessage(newErr.Error())
		newErr.LogError()
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	if !strings.Contains(res.Header.Get("Content-type"), "application/json") {
		newErr := exception.NewInternalErr(108, "Slack sent a response back but it should be application/json")
		errMsg := model.NewErrorMessage(newErr.Error())
		newErr.LogError()
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	body := new(model.OAuthResp)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		errMsg := model.NewErrorMessage(err.Error())
		env.Logger.Println(err.Error())
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	if !body.Ok {
		//TODO:replace with template!
		errMsg := model.NewErrorMessage("I am sorry we could not authorize you: " + body.Error)
		helper.Write(w, http.StatusUnauthorized, errMsg)
		return
	}
	//TODO:save token in db
	//TODO:replace with template!
	errMsg := model.NewErrorMessage("Added your team to the roster!")
	helper.Write(w, http.StatusOK, errMsg)
}
