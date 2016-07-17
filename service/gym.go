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
			return nil, exception.NewInternalError(err.Error())
		}
		gyms = append(gyms, gym)
	}
	if rows.Err() != nil {
		return nil, exception.NewInternalError(rows.Err().Error())
	}
	return gyms, nil
}

//AddGym addes a gym to the database
func AddGym(db *sql.DB, teamid string, gym *model.Gym) error {
	if len(gym.Name) > 50 {
		gym.Name = gym.Name[:50]
	}
	rows, err := db.Query("INSERT INTO system.Gym (TeamId,GymName,PokeTeam,UpdatedBy) VALUES ($1,$2,$3,$4)", teamid, gym.Name, gym.OwnerTeam, gym.UpdatedBy.ID)
	if err != nil {
		return exception.NewInternalError(err.Error())
	}
	defer rows.Close()
	return nil
}

//RemoveGym removes a gym from the watch list
func RemoveGym(db *sql.DB, teamid string, gymid int) (*model.Gym, error) {
	gym, err := GetGym(db, teamid, gymid)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("DELETE FROM system.Gym WHERE TeamId = $1 AND GymId = $2", teamid, gym.ID)
	if err != nil {
		return nil, exception.NewInternalError(err.Error())
	}
	defer rows.Close()
	return gym, nil
}

//GetGym gets a gym by its id number
func GetGym(db *sql.DB, teamid string, gymid int) (*model.Gym, error) {
	gym := new(model.Gym)
	gym.UpdatedBy = new(model.Trainer)
	err := db.QueryRow("SELECT GymId,GymName,PokeTeam,GymLevel,UpdatedBy,Updated FROM system.Gym WHERE TeamId = $1 AND GymId = $2", teamid, gymid).Scan(&gym.ID, &gym.Name, &gym.OwnerTeam, &gym.Level, &gym.UpdatedBy.ID, &gym.Updated)
	switch {
	case err == sql.ErrNoRows:
		return nil, exception.NewNoGymWithIDError()
	case err != nil:
		return nil, exception.NewInternalError(err.Error())
	default:
		return gym, nil
	}
}

//SplitGymsByTeam splits the gyms into a map by teams
func SplitGymsByTeam(gyms []*model.Gym) map[model.TeamEnum][]*model.Gym {
	splitGyms := make(map[model.TeamEnum][]*model.Gym)
	for _, gym := range gyms {
		splitGyms[gym.OwnerTeam] = append(splitGyms[gym.OwnerTeam], gym)
	}
	return splitGyms
}
