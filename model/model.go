package model

import "encoding/json"

//Publicer is an interface to state whether a model can be made to view publicly
type Publicer interface {
	Public() interface{}
}

//Jsonify turns the interface into a json object as a slice of bytes
func Jsonify(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}
