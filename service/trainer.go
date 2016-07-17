package service

import (
	"database/sql"

	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

//GetTrainer gets a trainer from the database
func GetTrainer(db *sql.DB, teamid string, trainerid string) (*model.Trainer, error) {
	trainer := new(model.Trainer)
	trainer.VerifiedBy = new(model.Trainer)
	err := db.QueryRow("SELECT TrainerId,UserName,VerifiedTeam,VerifiedBy,Updated FROM system.Trainer WHERE TeamId = $1 AND TrainerId = $2", teamid, trainerid).Scan(&trainer.ID, &trainer.UserName, &trainer.VerifiedTeam, &trainer.VerifiedBy.ID, &trainer.Updated)
	switch {
	case err == sql.ErrNoRows:
		return nil, exception.NewNoTrainerWitIDError()
	case err != nil:
		return nil, exception.NewInternalError(err.Error())
	default:
		return trainer, nil
	}
}

//InsertTrainer inserts a trainer into the database
func InsertTrainer(db *sql.DB, teamid string, trainer *model.Trainer) error {
	rows, err := db.Query("INSERT INTO system.Trainer (TrainerId,TeamId,UserName,VerifiedTeam) VALUES ($1,$2,$3,$4)", trainer.ID, teamid, trainer.UserName, trainer.VerifiedTeam)
	if err != nil {
		return exception.NewInternalError(err.Error())
	}
	defer rows.Close()
	return nil
}
