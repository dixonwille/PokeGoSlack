package handler

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/controller"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
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
				newErr := exception.NewOAuthAccessDeniedError()
				helper.WriteError(w, newErr)
				return
			}
			var msg string
			for _, v := range value {
				msg += msg + " " + v
			}
			newErr := exception.NewInternalError(msg)
			helper.WriteError(w, newErr)
			return
		}
	}
	code := params["code"]
	if code == nil {
		newErr := exception.NewInternalError("code was not part of the query")
		helper.WriteError(w, newErr)
		return
	}
	if code[0] == "" {
		newErr := exception.NewInternalError("Could not recieve the code from the request.")
		helper.WriteError(w, newErr)
		return
	}
	con, err := model.GetReqContext(r)
	if err != nil {
		helper.WriteError(w, err)
		return
	}
	con.OAuthCode = code[0]
	con.Set(r)
	controller.OAuthAccess(w, r)
}
