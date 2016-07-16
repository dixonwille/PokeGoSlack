package service

import (
	"database/sql"

	"github.com/davecgh/go-spew/spew"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

//InsertTeam inserts a new team into the database from response
func InsertTeam(db *sql.DB, body *model.OAuthResp) error {
	var teamname string
	err := db.QueryRow("SELECT TeamName FROM system.Team WHERE TeamId = $1", body.TeamID).Scan(&teamname)
	switch {
	case err == sql.ErrNoRows:
		if len(body.TeamName) > 50 {
			body.TeamName = body.TeamName[:50]
		}
		spew.Dump(len(body.AccessToken))
		rows, er := db.Query("INSERT INTO system.Team (TeamId,TeamName,AccessToken) VALUES ($1, $2, $3)", body.TeamID, body.TeamName, body.AccessToken)
		if er != nil {
			return er
		}
		defer rows.Close()
	case err != nil:
		return err
	default:
		return exception.ErrTeamExist
	}

	return nil
}
