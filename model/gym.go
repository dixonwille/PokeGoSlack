package model

//Gym is stats for the gym
type Gym struct {
	ID        int
	Name      string
	OwnerTeam TeamEnum
	Level     int
}
