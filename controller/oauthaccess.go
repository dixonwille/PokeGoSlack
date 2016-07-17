package controller

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/dixonwille/PokeGoSlack/service"
	"github.com/dixonwille/PokeGoSlack/slackapi"
)

//OAuthAccess is to ask for permission to connecet to slack
func OAuthAccess(w http.ResponseWriter, r *http.Request) {
	con, err := model.GetReqContext(r)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	if con.OAuthCode == "" {
		err = exception.NewInternalError("Could not get the OAuthCode from context")
		helper.WriteError(w, err)
		return
	}
	body, err := slackapi.OAuthAccess(con.OAuthCode)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	if !body.Ok {
		//TODO:replace with template that says goodbye
		msg := model.NewPrivateResponse("I am sorry we could not authorize you: " + body.Error)
		helper.Write(w, http.StatusOK, msg)
		return
	}

	if con.DB == nil {
		err = exception.NewInternalError("Could not get the database from context")
		helper.WriteError(w, err)
		return
	}
	team := &model.Team{
		ID:    body.TeamID,
		Name:  body.TeamName,
		Team:  model.None,
		Token: body.AccessToken,
	}
	sqlTeam, err := service.GetTeam(con.DB, team.ID)
	if err != nil && exception.IsNoTeamWithIDErr(err) {
		er := service.InsertTeam(con.DB, team)
		if er != nil {
			helper.WriteError(w, er)
			return
		}
	} else if err != nil {
		helper.WriteError(w, err)
		return
	}

	if sqlTeam != nil {
		err = exception.NewTeamExistError()
		helper.WriteError(w, err)
		return
	}

	//TODO:replace with template that asks for team color
	msg := model.NewPrivateResponse("Added your team to the roster!")
	helper.Write(w, http.StatusOK, msg)
}
