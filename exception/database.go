package exception

import (
	"errors"
	"net/http"
)

var (
	errTeamExist       = errors.New("This team is already using PokeGo.")
	errNoGymWithID     = errors.New("Could not find a gym with that ID.")
	errNoTrainerWithID = errors.New("Could not find a trainer with that ID.")
	errNoTeamWithID    = errors.New("Could not find a team with that ID.")
)

//NewTeamExistError is used to create an error when the team already exist
func NewTeamExistError() *Exception {
	return NewExceptionFromError(errTeamExist, http.StatusBadRequest)
}

//NewNoGymWithIDError is used if no gyms were found
func NewNoGymWithIDError() *Exception {
	return NewExceptionFromError(errNoGymWithID, http.StatusBadRequest)
}

//NewNoTrainerWitIDError is used if no gyms were found
func NewNoTrainerWitIDError() *Exception {
	return NewExceptionFromError(errNoTrainerWithID, http.StatusBadRequest)
}

//NewNoTeamWithIDError is used if not team was found
func NewNoTeamWithIDError() *Exception {
	return NewExceptionFromError(errNoTeamWithID, http.StatusBadRequest)
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

//IsNoTrainerWithIDErr whether gyms exists for this team
func IsNoTrainerWithIDErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errNoTrainerWithID.Error()
}

//IsNoTeamWithIDErr whether the error is because a team does not exist
func IsNoTeamWithIDErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errNoTeamWithID.Error()
}
