package controller

import "net/http"

var (
	emptyCommand Command = Command{}
)

//Command are commands used by and endpoint
type Command struct {
	Cmd        string
	HelpText   string
	Args       []Argument
	Controller func(http.ResponseWriter, *http.Request)
}

//Argument are the arguments that each command can have
type Argument struct {
	Name     string
	HelpText string
}

//NewCommand creates a new command to use
func NewCommand(cmd, help string) *Command {
	return &Command{
		Cmd:      cmd,
		HelpText: help,
	}
}

//AddArgument creates arguments for the commands
func (cmd *Command) AddArgument(name, help string) {
	var arg = &Argument{
		Name:     name,
		HelpText: help,
	}
	cmd.Args = append(cmd.Args, *arg)
}

//AddConroller adds a controller the command calls
func (cmd *Command) AddConroller(ctrl func(http.ResponseWriter, *http.Request)) {
	cmd.Controller = ctrl
}
