package exception

import (
	"errors"
	"net/http"
)

var (
	errTeamExist = errors.New("This team is already using PokeGo.")
)

//NewTeamExistError is used to create an error when the team already exist
func NewTeamExistError() *Exception {
	return NewExceptionFromError(errTeamExist, http.StatusBadRequest)
}

//IsTeamExistErr states whether this erro is a team exist error
func IsTeamExistErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errTeamExist.Error()
}
