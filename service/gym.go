package service

import (
	"database/sql"

	"github.com/dixonwille/PokeGoSlack/model"
)

//GetListGym gets a list of gyms and their statuses for a team
func GetListGym(db *sql.DB) ([]model.Gym, error) {
	return nil, nil
}
