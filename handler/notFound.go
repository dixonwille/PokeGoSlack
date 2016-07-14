package handler

import (
	"fmt"
	"net/http"

	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
)

//NotFound returns a 404 status with JSON message
func NotFound(w http.ResponseWriter, r *http.Request) {
	msg := model.NewErrorMessage(fmt.Sprintf("Endpoint not found: %s", r.RequestURI))
	helper.Write(w, http.StatusNotFound, msg)
}
