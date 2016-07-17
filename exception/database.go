package exception

import (
	"errors"
	"net/http"
)

var (
	errTeamExist     = errors.New("This team is already using PokeGo.")
	errNoGymsForTeam = errors.New("This team does not have any gyms.")
)

//NewTeamExistError is used to create an error when the team already exist
func NewTeamExistError() *Exception {
	return NewExceptionFromError(errTeamExist, http.StatusBadRequest)
}

//NewNoGymsForTeamError is used if no gyms were found
func NewNoGymsForTeamError() *Exception {
	return NewExceptionFromError(errNoGymsForTeam, http.StatusBadRequest)
}

//IsTeamExistErr states whether this erro is a team exist error
func IsTeamExistErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errTeamExist.Error()
}

//IsNoGymsForTeamErr whether gyms exists for this team
func IsNoGymsForTeamErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errNoGymsForTeam.Error()
}
