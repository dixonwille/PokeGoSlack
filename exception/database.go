package exception

import "errors"

var (
	//ErrTeamExist says that the team is already in DB
	ErrTeamExist = errors.New("This is already using PokeGo.")
)

//IsTeamExistErr states whether this erro is a team exist error
func IsTeamExistErr(err error) bool {
	return err.Error() == ErrTeamExist.Error()
}
