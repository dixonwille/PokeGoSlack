package adapter

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/model"
	"github.com/gorilla/context"
)

//InitContext initializes a context model to use
func InitContext() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqCon := new(model.ReqContext)
			context.Set(r, env.KeyReq, reqCon)
			h.ServeHTTP(w, r)
		})
	}
}
