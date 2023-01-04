package main

import (
	"errors"
)

var _ istorage = &storage{}

type storage struct {
	notes map[int]note
}

func newStorage() *storage {
	return &storage{
		notes: map[int]note{},
	}
}

func (s *storage) get(id int) *note {
	if n, ok := s.notes[id]; ok {
		return &n
	}

	return nil
}

func (s *storage) getAll() map[int]note {
	return s.notes
}

func (s *storage) add(n note) {
	s.notes[n.ID] = n
}

func (s *storage) delete(id int) error {
	if _, ok := s.notes[id]; ok {
		delete(s.notes, id)

		return nil
	}

	return errors.New("note not found")
}
