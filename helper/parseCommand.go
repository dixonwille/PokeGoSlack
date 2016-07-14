package helper

import (
	"strings"

	"github.com/dixonwille/PokeGoSlack/model"
)

//ParseCommand is used to parse out what the request wants to do
func ParseCommand(req model.Request) (string, []string) {
	text := req.Text
	strs := strings.Split(text, " ")
	if len(strs) == 1 {
		return strs[0], nil
	} else if len(strs) > 1 {
		return strs[0], strs[1:]
	} else {
		return "", nil
	}
}
