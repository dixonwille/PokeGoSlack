package model

//Team is used for Slack teams
type Team struct {
	ID    string
	Name  string
	Team  TeamEnum
	Token string
}
