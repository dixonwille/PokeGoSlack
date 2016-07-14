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
	router.HandleFunc("/", handler.Gym).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	middleRouter := http.NewServeMux()
	middleRouter.Handle("/", adapter.Adapt(router,
		adapter.Logging(env.Logger),
		adapter.Header("Content-Type", "application/json"),
	))
	fmt.Println("Listening on port " + env.Port)
	err := http.ListenAndServe(":"+env.Port, middleRouter)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
