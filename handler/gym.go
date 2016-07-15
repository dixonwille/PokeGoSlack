package handler

import (
	"net/http"
	"strings"

	"github.com/dixonwille/PokeGoSlack/controller"
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
	cmd, args := helper.ParseCommand(req)
	context.Set(r, env.KeyArgs, args)
	cmd = strings.ToLower(cmd)
	foundCtrl := false
	for _, command := range controller.GymCmds {
		if command.Cmd == cmd {
			foundCtrl = true
			context.Set(r, env.KeyCmd, command)
			command.Controller(w, r)
		}
	}
	if !foundCtrl {
		controller.GymHelp(w, r)
	}
}
