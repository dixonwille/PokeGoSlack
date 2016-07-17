package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/dixonwille/PokeGoSlack/service"
)

//GymCmds are the Possible commands to use with the gym endpoint
var GymCmds map[string]*model.Command

func gymInit() {
	GymCmds = make(map[string]*model.Command)
	addCmd := model.NewCommand("add", "Adds a new gym.")
	addCmd.AddConroller(AddGym)
	addCmd.AddArgument("gymname", "The name of the gym you are adding.")
	addCmd.AddArgument("help", "Displays this message.")
	GymCmds[addCmd.Cmd] = addCmd

	listCmd := model.NewCommand("list", "List all the gyms (with IDs) and the team that holds it.")
	listCmd.AddConroller(ListGyms)
	listCmd.AddArgument("gymid", "[OPTIONAL] See more information about the gym with that ID.")
	listCmd.AddArgument("help", "Displays this message.")
	GymCmds[listCmd.Cmd] = listCmd

	updateCmd := model.NewCommand("update", "Updates a gym's information.")
	updateCmd.AddConroller(UpdateGym)
	updateCmd.AddArgument("gymid", "The ID of the gym you want to update.")
	updateCmd.AddArgument("team", "Which team owns this gym.")
	updateCmd.AddArgument("level", "[OPTIONAL] What level is the gym.")
	updateCmd.AddArgument("help", "Displays this message.")
	GymCmds[updateCmd.Cmd] = updateCmd

	removeCmd := model.NewCommand("remove", "Removes a gym from listing (All data will be lost).")
	removeCmd.AddConroller(RemoveGym)
	removeCmd.AddArgument("gymid", "The ID of the gym you wish to remove.")
	removeCmd.AddArgument("help", "Displays this message.")
	GymCmds[removeCmd.Cmd] = removeCmd
}

//AddGym is used to add a gym to watch.
func AddGym(w http.ResponseWriter, con *model.ReqContext) {
	res := model.NewPrivateResponse("The command " + con.Command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//ListGyms is used to list all the gyms.
func ListGyms(w http.ResponseWriter, con *model.ReqContext) {
	imediateResp := model.RespondLater(true)
	helper.Write(w, http.StatusOK, imediateResp)
	if con.Args == nil || len(con.Args) == 0 {
		gyms, err := service.GetListGyms(con.DB, con.Form.TeamID)
		if err != nil {
			if exception.IsNoGymsForTeamErr(err) {
				res := model.NewPublicResponse("Your team is not watching any gyms! Use `/gym add` to start watching.")
				helper.RespondingLater(con.Form.ResponseURL, res)
				return
			}
			helper.RespondingLaterError(con.Form.ResponseURL, err)
			return
		}
		gymsResp := model.NewPublicResponse("Your team is watching the following:")
		splitGyms := service.SplitGymsByTeam(gyms)
		for tenum, gyms := range splitGyms {
			gymsAtt := model.NewAttachment("")
			gymsAtt.Text = model.PokeTeams[tenum].Name
			gymsAtt.Color = model.PokeTeams[tenum].Color
			for _, gym := range gyms {
				gymField := model.NewField(gym.Name, strconv.Itoa(gym.ID), true)
				gymsAtt.AddFields(*gymField)
			}
			gymsResp.AddAttachments(*gymsAtt)
		}
		helper.RespondingLater(con.Form.ResponseURL, gymsResp)
	}
	res := model.NewPrivateResponse("The command " + con.Command.Cmd + " has not been implimented yet")
	helper.RespondingLater(con.Form.ResponseURL, res)
}

//UpdateGym is used to update a specific gym.
func UpdateGym(w http.ResponseWriter, con *model.ReqContext) {
	res := model.NewPrivateResponse("The command " + con.Command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//RemoveGym removes a gym from the watch list.
func RemoveGym(w http.ResponseWriter, con *model.ReqContext) {
	res := model.NewPrivateResponse("The command " + con.Command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//GymHelp displays the help for the gym enpoint
func GymHelp(w http.ResponseWriter, helpCmd string) {
	if helpCmd != "" {
		res := cmdHelp(helpCmd)
		helper.Write(w, http.StatusOK, res)
		return
	}
	res := mainHelp()
	helper.Write(w, http.StatusOK, res)
}

func mainHelp() *model.Response {
	res := model.NewPrivateResponse("")
	att := model.NewAttachment("Help for `/gym`")
	att.Text = "Type help after a command to get more information"
	for _, cmd := range GymCmds {
		title := "/gym " + cmd.Cmd
		field := model.NewField(title, cmd.HelpText, false)
		att.AddFields(*field)
	}
	res.AddAttachments(*att)
	return res
}

func cmdHelp(cmdHelp string) *model.Response {
	command := GymCmds[cmdHelp]
	res := model.NewPrivateResponse("")
	att := model.NewAttachment("Help for `/gym " + command.Cmd + "`")
	att.Title = "/gym " + command.Cmd
	att.Text = command.HelpText
	for i, arg := range command.Args {
		if i != len(command.Args)-1 {
			att.Title = att.Title + " <" + arg.Name + ">"
			title := "<" + strings.ToUpper(arg.Name) + ">"
			field := model.NewField(title, arg.HelpText, true)
			att.AddFields(*field)
		}
	}
	res.AddAttachments(*att)
	return res
}
