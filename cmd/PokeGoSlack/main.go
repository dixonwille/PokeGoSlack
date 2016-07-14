package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dixonwille/PokeGoSlack/adapter"
	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/gym", handler.Gym).Methods("POST")
	router.HandleFunc("/trainer", handler.Trainer).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	middleRouter := http.NewServeMux()
	middleRouter.Handle("/", adapter.Adapt(router,
		adapter.Logging(env.Logger),
		adapter.Header("Content-Type", "application/json"),
	))
	middleRouter.Handle("/gym", adapter.Adapt(router,
		adapter.Validate("/gym"),
	))
	fmt.Println("Listening on port " + env.Port)
	err := http.ListenAndServe(":"+env.Port, middleRouter)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
