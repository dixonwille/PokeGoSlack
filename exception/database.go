package exception

import (
	"errors"
	"net/http"
)

var (
	errTeamExist   = errors.New("This team is already using PokeGo.")
	errNoGymWithID = errors.New("Could not find a gym with that ID.")
)

//NewTeamExistError is used to create an error when the team already exist
func NewTeamExistError() *Exception {
	return NewExceptionFromError(errTeamExist, http.StatusBadRequest)
}

//NewNoGymWithIDError is used if no gyms were found
func NewNoGymWithIDError() *Exception {
	return NewExceptionFromError(errNoGymWithID, http.StatusBadRequest)
}

//IsTeamExistErr states whether this erro is a team exist error
func IsTeamExistErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errTeamExist.Error()
}

//IsNoGymWithIDErr whether gyms exists for this team
func IsNoGymWithIDErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errNoGymWithID.Error()
}
