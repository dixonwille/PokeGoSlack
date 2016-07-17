package model

import (
	"database/sql"
	"time"
)

//Trainer is the pokemon trainer
type Trainer struct {
	ID           sql.NullString
	UserName     string
	VerifiedTeam TeamEnum
	VerifiedBy   *Trainer
	Updated      time.Time
}
