package model

import (
	"encoding/json"

	"github.com/dixonwille/PokeGoSlack/exception"
)

func init() {
	PokeTeams = make(map[TeamEnum]PokeTeam)
	PokeTeams[None] = PokeTeam{
		Name:  "Available",
		Color: "#D3D3D3",
	}
	PokeTeams[Mystic] = PokeTeam{
		Name:  "Mystic",
		Color: "#1977F6",
	}
	PokeTeams[Valor] = PokeTeam{
		Name:  "Valor",
		Color: "#EF1600",
	}
	PokeTeams[Instinct] = PokeTeam{
		Name:  "Instinct",
		Color: "#FDD100",
	}
}

//Publicer is an interface to state whether a model can be made to view publicly
type Publicer interface {
	Public() interface{}
}

//Jsonify turns the interface into a json object as a slice of bytes
func Jsonify(i interface{}) ([]byte, error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return nil, exception.NewInternalError(err.Error())
	}
	return bytes, nil
}
