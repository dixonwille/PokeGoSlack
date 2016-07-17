package model

import "time"

//Trainer is the pokemon trainer
type Trainer struct {
	ID           string
	UserName     string
	VerifiedTeam TeamEnum
	VerifiedBy   *Trainer
	Updated      time.Time
}
