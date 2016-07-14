package adapter

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/gorilla/context"
)

//Validate makes sure that the request on endpoint is valid
func Validate(command string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := helper.ValidateRequestAndParse(r, command)
			if req == nil {
				return //error was handled for us already
			}
			context.Set(r, env.KeyForm, req)
			h.ServeHTTP(w, r)
		})
	}
}
