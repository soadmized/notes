package main

type iservice interface {
	getNotes() map[int]note
	getNoteByID(id int) *note
	addNote(n note)
	// ...
}
