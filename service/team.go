package service

import (
	"database/sql"

	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

//InsertTeam inserts a new team into the database from response
func InsertTeam(db *sql.DB, body *model.OAuthResp) error {
	var teamname string
	err := db.QueryRow("SELECT TeamName FROM system.Team WHERE TeamId = $1", body.TeamID).Scan(&teamname)
	switch {
	case err == sql.ErrNoRows:
		_, err = db.Query("INSERT INTO system.Team (TeamId,TeamName,AccessToken) VALUES ($1, $2, $3)", body.TeamID, body.TeamName, body.AccessToken)
		if err != nil {
			return err
		}
	case err != nil:
		return err
	default:
		return exception.ErrTeamExist
	}

	return nil
}
