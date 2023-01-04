package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type iapi interface {
	run()
}

type api struct {
	router  *mux.Router
	service iservice
}

func newAPI() *api {
	a := api{
		router:  mux.NewRouter(),
		service: newService(),
	}

	a.registryHandlers()

	return &a
}

func (a *api) registryHandlers() {
	a.router.HandleFunc("/note", a.getNotesSteamHandleFunc()).Methods(http.MethodGet)
	a.router.HandleFunc("/note/{id}", a.getNoteByIDSteamHandleFunc()).Methods(http.MethodGet)
	a.router.HandleFunc("/note", a.addNoteSteamHandleFunc()).Methods(http.MethodPost)
}

func (a *api) getNotesSteamHandleFunc() http.HandlerFunc {
	return func(writer http.ResponseWriter, _ *http.Request) {
		enc := json.NewEncoder(writer)
		if err := enc.Encode(a.service.getNotes()); err != nil {
			log.Println(err)
		}
	}
}

func (a *api) getNoteByIDSteamHandleFunc() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		enc := json.NewEncoder(writer)
		if err := enc.Encode(a.service.getNoteByID(id)); err != nil {
			log.Println(err)
		}
	}
}

func (a *api) addNoteSteamHandleFunc() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		n := note{}

		if err := dec.Decode(&n); err != nil {
			log.Println(err)
			return
		}

		a.service.addNote(n)
	}
}

func (a *api) run() {
	log.Fatal(http.ListenAndServe(":8080", a.router))
}
