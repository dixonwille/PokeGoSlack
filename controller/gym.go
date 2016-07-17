package controller

import (
	"database/sql"
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
	if con.Args == nil || len(con.Args) == 0 {
		res := cmdHelp(con.Command.Cmd)
		helper.Write(w, http.StatusOK, res)
		return
	}
	gymName := strings.Join(con.Args, " ")
	gym := &model.Gym{
		Name:      gymName,
		OwnerTeam: model.None,
		UpdatedBy: &model.Trainer{
			ID:           sql.NullString{String: con.Form.UserID, Valid: true},
			UserName:     con.Form.UserName,
			VerifiedTeam: model.None,
		},
	}
	err := service.AddGym(con.DB, con.Form.TeamID, gym)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	res := model.NewPublicResponse("The gym *" + gymName + "* was added to your team.")
	helper.Write(w, http.StatusOK, res)
}

//ListGyms is used to list all the gyms.
func ListGyms(w http.ResponseWriter, con *model.ReqContext) {
	if con.Args != nil && len(con.Args) > 1 {
		res := cmdHelp(con.Command.Cmd)
		helper.Write(w, http.StatusOK, res)
		return
	}
	if con.Args == nil || len(con.Args) == 0 {
		gyms, err := service.GetListGyms(con.DB, con.Form.TeamID)
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		if len(gyms) == 0 {
			res := model.NewPublicResponse("Your team is not watching any gyms! Use `/gym add` to start watching.")
			helper.Write(w, http.StatusOK, res)
			return
		}
		gymsResp := model.NewPublicResponse("Your team is watching the following:")
		splitGyms := service.SplitGymsByTeam(gyms)
		for tenum, gyms := range splitGyms {
			gymsAtt := model.NewAttachment("")
			gymsAtt.Title = model.PokeTeams[tenum].Name
			if tenum == model.None {
				gymsAtt.Text = "These gyms are available for anyone to capture."
			} else {
				gymsAtt.Text = "These gyms are controlled by team " + model.PokeTeams[tenum].Name
			}
			gymsAtt.Color = model.PokeTeams[tenum].Color
			for _, gym := range gyms {
				gymField := model.NewField(gym.Name, "ID: "+strconv.Itoa(gym.ID), true)
				gymsAtt.AddFields(*gymField)
			}
			gymsResp.AddAttachments(*gymsAtt)
		}
		helper.Write(w, http.StatusOK, gymsResp)
		return
	}
	gymid, err := strconv.Atoi(con.Args[0])
	if err != nil {
		newErr := exception.NewNotANumberError()
		helper.WriteError(w, newErr)
		return
	}
	gym, err := service.GetGym(con.DB, con.Form.TeamID, gymid)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	gymRes := model.NewPublicResponse("")
	gymAtt := model.NewAttachment("")
	gymAtt.Title = gym.Name
	gymAtt.Text = "ID: " + strconv.Itoa(gym.ID)
	gymAtt.Color = model.PokeTeams[gym.OwnerTeam].Color
	teamField := model.NewField("Team", model.PokeTeams[gym.OwnerTeam].Name, true)
	levelField := model.NewField("Level", strconv.Itoa(gym.Level), true)
	gymAtt.AddFields(*teamField, *levelField)
	gymRes.AddAttachments(*gymAtt)
	helper.Write(w, http.StatusOK, gymRes)
}

//UpdateGym is used to update a specific gym.
func UpdateGym(w http.ResponseWriter, con *model.ReqContext) {
	res := model.NewPrivateResponse("The command " + con.Command.Cmd + " has not been implimented yet")
	helper.Write(w, http.StatusOK, res)
}

//RemoveGym removes a gym from the watch list.
func RemoveGym(w http.ResponseWriter, con *model.ReqContext) {
	if con.Args == nil || len(con.Args) != 1 {
		res := cmdHelp(con.Command.Cmd)
		helper.Write(w, http.StatusOK, res)
		return
	}
	gymID, err := strconv.Atoi(con.Args[0])
	if err != nil {
		newErr := exception.NewNotANumberError()
		helper.WriteError(w, newErr)
		return
	}
	gym, err := service.RemoveGym(con.DB, con.Form.TeamID, gymID)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	res := model.NewPublicResponse("*" + gym.Name + "* was removed from the watch list.")
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
