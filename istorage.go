package main

type istorage interface {
	get(id int) *note
	getAll() map[int]note
	add(n note)
	delete(id int) error
}
