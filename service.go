package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var _ ServiceInterface = &service{}

type service struct {
	repo *storage
}

type ServiceInterface interface {
	getNotes(w http.ResponseWriter, r *http.Request)
	createNote(w http.ResponseWriter, r *http.Request)
	getNoteByID(w http.ResponseWriter, r *http.Request)
	deleteNoteByID(w http.ResponseWriter, r *http.Request)
	updateNoteByID(w http.ResponseWriter, r *http.Request)
}

func (s *service) getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	notes := s.repo.getAll()
	err := json.NewEncoder(w).Encode(notes)

	if err != nil {
		return
	}
}

func (s *service) createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	n := note{}

	n.ID = rand.Intn(999 - 10 + 1)
	_ = json.NewDecoder(r.Body).Decode(&n)
	s.repo.add(n)

	err := json.NewEncoder(w).Encode(n)
	if err != nil {
		return
	}
}

func (s *service) getNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if n := s.repo.get(id); n != nil {
		err := json.NewEncoder(w).Encode(n)
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(&note{}) // return empty note
		if err != nil {
			return
		}
	}
}

func (s *service) deleteNoteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := s.repo.delete(id); err == nil {
		e := json.NewEncoder(w).Encode("Note was deleted")

		if e != nil {
			return
		}
	} else {
		e := json.NewEncoder(w).Encode("Note was not found")

		if e != nil {
			return
		}
	}
}

func (s *service) updateNoteByID(w http.ResponseWriter, r *http.Request) {
	s.repo.update()
}
