package adapter

import (
	"database/sql"
	"net/http"

	"github.com/dixonwille/PokeGoSlack/helper"
	"github.com/dixonwille/PokeGoSlack/model"
)

//Database adds a database instance to the conext.
func Database(db *sql.DB) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			con, err := model.GetReqContext(r)
			if err != nil {
				helper.WriteError(w, err)
				return
			}
			con.DB = db
			con.Set(r)
			h.ServeHTTP(w, r)
		})
	}
}
