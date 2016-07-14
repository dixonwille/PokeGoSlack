package handler

import (
	"net/http"
	"strings"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/context"
)

//Gym handles all request comming in.
//Parses the form and directs to controllers.
func Gym(w http.ResponseWriter, r *http.Request) {
	req, ok := context.Get(r, env.KeyForm).(model.Request)
	if !ok {
		newError := exception.NewInternalErr(101, "Could not get request object")
		errMsg := model.NewErrorMessage(newError.Error())
		newError.LogError()
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	cmd, _ := helper.ParseCommand(req)
	context.Set(r, env.KeyCMD, cmd)
	var res *model.Response
	switch strings.ToLower(cmd) {
	case "private":
		res = model.NewPrivateResponse("Hey there from PokeGoSlack API. You are the only one to see this.")
	case "public":
		res = model.NewPublicResponse("Hey there from PokeGoSlack API. Everyone is able to see this.")
	default:
		context.Set(r, env.KeyCMD, "help")
		res = helpResponse()
	}
	helper.Write(w, http.StatusOK, res)
}

func helpResponse() *model.Response {
	res := model.NewPrivateResponse("")
	priv := model.NewField("/gym private", "API will only respond to you.", false)
	pub := model.NewField("/gym public", "API will respond to everyone in channel.", false)
	hlp := model.NewField("/gym help", "Displays this message", false)
	att := model.NewAttachment("Possible commands for `/gym`")
	att.AddFields(*priv, *pub, *hlp)
	res.AddAttachments(*att)
	return res
}
