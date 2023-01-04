package main

var _ iservice = &newservice{}

type newservice struct {
	repo istorage
}

func newService() *newservice {
	return &newservice{
		repo: newStorage(),
	}
}

func (s *newservice) getNotes() map[int]note {
	return s.repo.getAll()
}

func (s *newservice) getNoteByID(id int) *note {
	return s.repo.get(id)
}

func (s *newservice) addNote(n note) {
	s.repo.add(n)
}
