package adapter

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
)

//Validate makes sure that the request on endpoint is valid
func Validate(command string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := helper.ParseAndValidateRequest(r)
			if err != nil {
				helper.WriteError(w, err)
				return
			}
			con, err := model.GetReqContext(r)
			if err != nil {
				helper.WriteError(w, err)
				return
			}
			con.Form = req
			con.Set(r)
			h.ServeHTTP(w, r)
		})
	}
}
