package service

import (
	"database/sql"

	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

//GetListGyms gets a list of gyms and their statuses for a team
func GetListGyms(db *sql.DB, teamid string) ([]*model.Gym, error) {
	rows, err := db.Query("SELECT GymId, GymName, PokeTeam, GymLevel FROM system.Gym WHERE TeamId = $1 ORDER BY PokeTeam", teamid)
	if err != nil {
		return nil, exception.NewInternalError(err.Error())
	}
	defer rows.Close()
	gyms := []*model.Gym{}
	for rows.Next() {
		gym := new(model.Gym)
		err := rows.Scan(&gym.ID, &gym.Name, &gym.OwnerTeam, &gym.Level)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, exception.NewNoGymsForTeamError()
			}
			return nil, exception.NewInternalError(err.Error())
		}
		gyms = append(gyms, gym)
	}
	if rows.Err() != nil {
		return nil, exception.NewInternalError(rows.Err().Error())
	}
	return gyms, nil
}

//SplitGymsByTeam splits the gyms into a map by teams
func SplitGymsByTeam(gyms []*model.Gym) map[model.TeamEnum][]*model.Gym {
	splitGyms := make(map[model.TeamEnum][]*model.Gym)
	for _, gym := range gyms {
		splitGyms[gym.OwnerTeam] = append(splitGyms[gym.OwnerTeam], gym)
	}
	return splitGyms
}
