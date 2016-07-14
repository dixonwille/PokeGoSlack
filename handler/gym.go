package handler

import (
	"net/http"
	"strings"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/context"
)

//Gym handles all request comming in.
//Parses the form and directs to controllers.
func Gym(w http.ResponseWriter, r *http.Request) {
	req, ok := context.Get(r, env.KeyForm).(model.Request)
	if !ok {
		return
	}
	cmd, _ := helper.ParseCommand(req)
	var res *model.Response
	switch strings.ToLower(cmd) {
	case "private":
		res = model.NewPrivateResponse("Hey there from PokeGoSlack API. You are the only one to see this.")
	case "public":
		res = model.NewPublicResponse("Hey there from PokeGoSlack API. Everyone is able to see this.")
	default:

	}
	helper.Write(w, http.StatusOK, res)
}

func helpResponse() *model.Response {
	res := model.NewPrivateResponse("")
	priv := model.NewField("/gym private", "API will only respond to you.", false)
	pub := model.NewField("/gym public", "API will respond to everyone in channel.", false)
	att := model.NewAttachment("Possible commands for `/gym`")
	att.AddFields(*priv, *pub)
	res.AddAttachments(*att)
	return res
}
