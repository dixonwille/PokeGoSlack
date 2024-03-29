package model

//TeamEnum are possible teams to be on
type TeamEnum int

//PokeTeams is the string representation of the enum
var PokeTeams map[TeamEnum]PokeTeam

const (
	//None is for no one owns this gym yet
	None TeamEnum = iota
	//Mystic is for team blue
	Mystic
	//Valor is for team red
	Valor
	//Instinct is for team yellow
	Instinct
)

//PokeTeam are the attributes associated with each team
type PokeTeam struct {
	Name   string
	Color  string
	Search string
}
