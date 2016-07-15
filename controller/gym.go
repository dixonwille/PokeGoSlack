package controller

import (
	"net/http"
	"strings"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/context"
)

//GymCmds are the Possible commands to use with the gym endpoint
var GymCmds []Command

func gymInit() {
	addCmd := NewCommand("add", "Adds a new gym.")
	addCmd.AddConroller(AddGym)
	addCmd.AddArgument("gymname", "The name of the gym you are adding.")
	addCmd.AddArgument("help", "Displays this message.")
	listCmd := NewCommand("list", "List all the gyms (with IDs) and the team that holds it.")
	listCmd.AddConroller(ListGyms)
	listCmd.AddArgument("gymid", "[OPTIONAL] See more information about the gym with that ID.")
	listCmd.AddArgument("help", "Displays this message.")
	updateCmd := NewCommand("update", "Updates a gym's information.")
	updateCmd.AddConroller(UpdateGym)
	updateCmd.AddArgument("gymid", "The ID of the gym you want to update.")
	updateCmd.AddArgument("team", "Which team owns this gym.")
	updateCmd.AddArgument("level", "[OPTIONAL] What level is they gym.")
	updateCmd.AddArgument("help", "Displays this message.")
	removeCmd := NewCommand("remove", "Removes a gym from listing (All data will be lost).")
	removeCmd.AddConroller(RemoveGym)
	removeCmd.AddArgument("gymid", "The ID of the gym you wish to remove.")
	removeCmd.AddArgument("help", "Displays this message.")
	GymCmds = []Command{*addCmd, *listCmd, *updateCmd, *removeCmd}
}

//AddGym is used to add a gym to watch.
func AddGym(w http.ResponseWriter, r *http.Request) {
	command, _ := parseReqAndCheckForHelp(w, r)
	res := model.NewPrivateResponse("The command " + command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//ListGyms is used to list all the gyms.
func ListGyms(w http.ResponseWriter, r *http.Request) {
	command, _ := parseReqAndCheckForHelp(w, r)
	res := model.NewPrivateResponse("The command " + command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//UpdateGym is used to update a specific gym.
func UpdateGym(w http.ResponseWriter, r *http.Request) {
	command, _ := parseReqAndCheckForHelp(w, r)
	res := model.NewPrivateResponse("The command " + command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//RemoveGym removes a gym from the watch list.
func RemoveGym(w http.ResponseWriter, r *http.Request) {
	command, _ := parseReqAndCheckForHelp(w, r)
	res := model.NewPrivateResponse("The command " + command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//GymHelp displays the help for the gym enpoint
func GymHelp(w http.ResponseWriter, r *http.Request) {
	helpCmd, ok := context.Get(r, env.KeyHelpCmd).(string)
	if ok && helpCmd != "" {
		res := cmdHelp(helpCmd)
		helper.Write(w, http.StatusOK, res)
		return
	}
	res := mainHelp()
	helper.Write(w, http.StatusOK, res)
}

func parseReqAndCheckForHelp(w http.ResponseWriter, r *http.Request) (Command, []string) {
	command, ok := context.Get(r, env.KeyCmd).(Command)
	if !ok {
		newErr := exception.NewInternalErr(103, "Could not get the command that was called")
		res := model.NewErrorMessage(newErr.Error())
		newErr.LogError()
		helper.Write(w, http.StatusInternalServerError, res)
		return Command{}, nil
	}
	args, ok := context.Get(r, env.KeyArgs).([]string)
	if !ok {
		newErr := exception.NewInternalErr(102, "Could not get arguments for "+command.Cmd)
		res := model.NewErrorMessage(newErr.Error())
		newErr.LogError()
		helper.Write(w, http.StatusInternalServerError, res)
		return Command{}, nil
	}
	switch len(args) {
	case 1:
		if args[0] == command.Args[len(command.Args)-1].Name {
			context.Set(r, env.KeyHelpCmd, command.Cmd)
			GymHelp(w, r)
			return Command{}, nil
		}
		if len(command.Args)-1 == 1 && command.Args[0].Name != args[0] {
			res := model.NewErrorMessage("Could not find the command " + args[0] + " for " + command.Cmd)
			helper.Write(w, http.StatusBadRequest, res)
			return Command{}, nil
		}
	case 0:
		context.Set(r, env.KeyHelpCmd, command.Cmd)
		GymHelp(w, r)
		return Command{}, nil
	default:
		for _, arg := range args {
			if !argExist(command, arg) {
				res := model.NewErrorMessage("Could not find the command " + arg + " for " + command.Cmd)
				helper.Write(w, http.StatusBadRequest, res)
				return Command{}, nil
			}
		}
	}
	return command, args
}

func mainHelp() *model.Response {
	res := model.NewPrivateResponse("")
	att := model.NewAttachment("Help for `/gym`")
	att.Title = "/gym"
	att.Text = "Used to update, list, add, and remove gyms.\nFollow any command with help to see more about that command."
	for _, cmd := range GymCmds {
		title := att.Title + cmd.Cmd
		field := model.NewField(title, cmd.HelpText, false)
		att.AddFields(*field)
	}
	res.AddAttachments(*att)
	return res
}

func cmdHelp(cmdHelp string) *model.Response {
	var command Command
	foundCmd := false
	for _, cmd := range GymCmds {
		if cmd.Cmd == cmdHelp {
			foundCmd = true
			command = cmd
		}
	}
	if !foundCmd {
		model.NewErrorMessage("Could not find the command " + cmdHelp + " for the gym endpoint")
	}
	res := model.NewPrivateResponse("")
	att := model.NewAttachment("Help for `/gym " + command.Cmd + "`")
	att.Title = "/gym " + command.Cmd
	att.Text = command.HelpText
	for _, arg := range command.Args {
		title := att.Title + " &lt;" + strings.ToUpper(arg.Name) + "&gt;"
		field := model.NewField(title, arg.HelpText, false)
		att.AddFields(*field)
	}
	res.AddAttachments(*att)
	return res
}

func argExist(cmd Command, arg string) bool {
	for _, cmdArg := range cmd.Args {
		if arg == cmdArg.Name {
			return true
		}
	}
	return false
}
