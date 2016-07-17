package model

import (
	"database/sql"
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/gorilla/context"
)

//ReqContext holds all information gathered
type ReqContext struct {
	Command   *Command
	Args      []string
	Form      *Request
	DB        *sql.DB
	OAuthCode string
}

//GetReqContext gets the context to use
func GetReqContext(r *http.Request) (*ReqContext, error) {
	con, ok := context.Get(r, env.KeyReq).(*ReqContext)
	if !ok {
		return nil, exception.NewInternalError("Could not get context from request.")
	}
	return con, nil
}

//Set sets the request context
func (con *ReqContext) Set(r *http.Request) {
	context.Set(r, env.KeyReq, con)
}
