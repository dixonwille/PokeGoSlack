package handler

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/controller"
	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/context"
)

const (
	accessDenied string = "access_denied"
)

//OAuth is handling new teams using the application
func OAuth(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	for param, value := range params {
		if param == "error" {
			if value[0] == accessDenied {
				//TODO:replace with sorry you don't want template
				errMsg := model.NewErrorMessage("Sorry something went wrong.")
				helper.Write(w, http.StatusOK, errMsg)
				return
			}
			var msg string
			for _, v := range value {
				msg += msg + " " + v
			}
			newErr := exception.NewInternalErr(104, msg)
			errMsg := model.NewErrorMessage(newErr.Error())
			newErr.LogError()
			helper.Write(w, http.StatusInternalServerError, errMsg)
			return
		}
	}
	code := params["code"][0]
	if code == "" {
		newErr := exception.NewInternalErr(105, "Could not recieve the code from the request.")
		errMsg := model.NewErrorMessage(newErr.Error())
		newErr.LogError()
		helper.Write(w, http.StatusInternalServerError, errMsg)
		return
	}
	context.Set(r, env.KeyCode, code)
	controller.OAuthAccess(w, r)
}
