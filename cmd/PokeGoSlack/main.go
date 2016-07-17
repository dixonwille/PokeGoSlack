package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dixonwille/PokeGoSlack/adapter"
	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", env.DBConnString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(15) //20 Is the limit for heroku
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/gym", adapter.Adapt(http.HandlerFunc(handler.Gym),
		adapter.Database(db),
		adapter.Validate(),
		adapter.InitContext(),
		adapter.Header("Content-type", "application/json"),
		adapter.Logging(env.Logger),
	)).Methods("POST")

	router.Handle("/trainer", adapter.Adapt(http.HandlerFunc(handler.Trainer),
		adapter.Database(db),
		adapter.Validate(),
		adapter.InitContext(),
		adapter.Header("Content-type", "application/json"),
		adapter.Logging(env.Logger),
	)).Methods("POST")

	router.Handle("/oauth", adapter.Adapt(http.HandlerFunc(handler.OAuth),
		adapter.Database(db),
		adapter.InitContext(),
		adapter.Logging(env.Logger),
	)).Methods("GET")

	// router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	fmt.Println("Listening on port " + env.Port)
	err = http.ListenAndServe(":"+env.Port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
