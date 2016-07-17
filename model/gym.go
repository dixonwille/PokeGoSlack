package model

import "time"

//Gym is stats for the gym
type Gym struct {
	ID        int
	Name      string
	OwnerTeam TeamEnum
	Level     int
	UpdatedBy *Trainer
	Updated   time.Time
}
