package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func router(s ServiceInterface) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", greeting).Methods("GET")
	r.HandleFunc("/get", s.getNotes).Methods("POST")
	r.HandleFunc("/get/{id}", s.getNoteByID).Methods("POST")
	r.HandleFunc("/create", s.createNote).Methods("POST")
	r.HandleFunc("/delete/{id}", s.deleteNoteByID).Methods("POST")
	r.HandleFunc("/update/{id}", s.updateNoteByID).Methods("POST")

	return r
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
POST /create - makes a note
POST /get - get all notes
POST /get/:id - get note with id
POST /delete/:id - delete note with id`)
}
