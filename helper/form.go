package helper

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/schema"
)

//ParseForm is used to parse a request form
func ParseForm(r *http.Request) *model.Request {
	err := r.ParseForm()
	if err != nil {
		newError := exception.NewInternalErr(200, err.Error())
		errMsg := model.NewErrorMessage("Could not parse the request form")
		newError.LogError()
		helper.Write(w, http.StatusBadRequest, errMsg)
		return nil
	}
	req := new(model.Request)
	dec := schema.NewDecoder()
	err = dec.Decode(req, r.PostForm)
	if err != nil {
		newError := exception.NewInternalErr(201, err.Error())
		errMsg := model.NewErrorMessage("Could not parse the request form")
		newError.LogError()
		helper.Write(w, http.StatusBadRequest, errMsg)
		return nil
	}
	return req
}

//ValidateRequestAndParse parses the request then validates.
//Returns the request form if it is valid.
func ValidateRequestAndParse(r *http.Request, command string) *model.Request {
	req := ParseForm(r)
	if req == nil {
		return nil //Error already written
	}

	if req.Token != env.Token {
		errMsg := model.NewErrorMessage("Request not accepted")
		helper.Write(w, http.StatusForbidden, errMsg)
		return nil //Invalid Request
	}

	if req.Commad != command {
		errMsg := model.NewErrorMessage("Unexpected Command")
		helper.Write(w, http.StatusBadRequest, errMsg)
		return nil //Invalid Request
	}
	return req
}
