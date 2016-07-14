package handler

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/helper"
)

type key int

const (
	errorKey key = iota
)

//Gym handles all request comming in.
//Parses the form and directs to controllers.
func Gym(w http.ResponseWriter, r *http.Request) {
	helper.Write(w, http.StatusOK, "")
}
