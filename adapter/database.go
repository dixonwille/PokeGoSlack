package adapter

import (
	"database/sql"
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/gorilla/context"
)

//Database adds a database instance to the conext.
func Database(db *sql.DB) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, env.KeyDB, db)
			h.ServeHTTP(w, r)
		})
	}
}
