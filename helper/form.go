package helper

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/schema"
)

//ParseForm is used to parse a request form
func ParseForm(r *http.Request) (*model.Request, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, exception.NewParseSlackRequestError()
	}
	req := new(model.Request)
	dec := schema.NewDecoder()
	err = dec.Decode(req, r.PostForm)
	if err != nil {
		return nil, exception.NewParseSlackRequestError()
	}
	return req, nil
}

//ParseAndValidateRequest parses the request then validates.
//Returns the request form if it is valid.
func ParseAndValidateRequest(r *http.Request) (*model.Request, error) {
	req, err := ParseForm(r)
	if err != nil {
		return nil, err
	}

	if req.Token != env.Token {
		return nil, exception.NewInvalidTokenError()
	}
	return req, nil
}
