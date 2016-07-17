package service

import (
	"database/sql"

	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

//InsertTeam inserts a new team into the database from response
func InsertTeam(db *sql.DB, team *model.Team) error {
	if len(team.Name) > 100 {
		team.Name = team.Name[:100]
	}
	rows, err := db.Query("INSERT INTO system.Team (TeamId,TeamName,PokeTeam,AccessToken) VALUES ($1, $2, $3, $4)", team.ID, team.Name, team.Team, team.Token)
	if err != nil {
		return exception.NewInternalError(err.Error())
	}
	defer rows.Close()
	return nil
}

//GetTeam gets the Team with id
func GetTeam(db *sql.DB, teamid string) (*model.Team, error) {
	team := new(model.Team)
	err := db.QueryRow("SELECT TeamId,TeamName,PokeTeam,AccessToken FROM system.Team WHERE TeamId = $1", teamid).Scan(&team.ID, &team.Name, &team.Team, &team.Token)
	switch {
	case err == sql.ErrNoRows:
		return nil, exception.NewNoTeamWithIDError()
	case err != nil:
		return nil, exception.NewInternalError(err.Error())
	default:
		return team, nil
	}
}
