package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type note struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Body string `json:"body"`
}

var (
	notes = []note{
		{
			ID: 1,
			Title: "First note",
			Author: "Alex",
			Body: "This is the very first note to check /get and /get/id endpoints",
		},
	}
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", greeting).Methods("GET")
	r.HandleFunc("/get", getNotes).Methods("POST")
	r.HandleFunc("/get/{id}", getNoteByID).Methods("POST")
	r.HandleFunc("/create", createNote).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteNoteByID).Methods("POST")
	r.HandleFunc("/update/{id}", updateNoteByID).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
POST /create - makes a note
POST /get - get all notes
POST /get/:id - get note with id
POST /delete/:id - delete note with id`)
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		return
	}
}

func getNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range notes {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&note{})
}

func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	note := note{}
	_ = json.NewDecoder(r.Body).Decode(&note)
	note.ID = rand.Intn(999 - 10 + 1)
	notes = append(notes, note)
	json.NewEncoder(w).Encode(note)
}

func deleteNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range notes {
		if item.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			json.NewEncoder(w).Encode("Note was deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Note was not found")
}

func updateNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range notes {
		if item.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			note := note{}
			_ = json.NewDecoder(r.Body).Decode(&note)
			note.ID = id
			notes = append(notes, note)
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode(notes)
}
