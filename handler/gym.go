package handler

import (
	"net/http"
	"strings"

	"github.com/dixonwille/PokeGoSlack/controller"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
)

//Gym handles all request comming in.
//Parses the form and directs to controllers.
func Gym(w http.ResponseWriter, r *http.Request) {
	con, err := model.GetReqContext(r)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	if con.Form == nil {
		err = exception.NewInternalError("Could not get the form from context")
		helper.WriteError(w, err)
		return
	}
	cmd, args := helper.ParseCommand(con.Form)
	if cmd == "" || cmd == "help" {
		controller.GymHelp(w, "")
		return
	}
	cmd = strings.ToLower(cmd)
	command := controller.GymCmds[cmd]
	if command == nil {
		err = exception.NewCmdNotFoundError()
		helper.WriteError(w, err)
		return
	}
	if len(args) > 0 && args[len(args)-1] == "help" {
		controller.GymHelp(w, command.Cmd)
		return
	}
	con.Command = command
	con.Args = args
	con.Set(r)
	command.Controller(w, con)
}
