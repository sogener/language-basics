package main

import "errors"

type programmer struct {
	id        int
	name      string
	role      string
	isWorking bool
}

type outSource interface {
	insert(p programmer) error
	get(id int) (p programmer, err error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]programmer
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{data: make(map[int]programmer)}
}

func (s *memoryStorage) insert(p programmer) error {
	s.data[p.id] = p
	return nil
}

func (s *memoryStorage) get(id int) (p programmer, err error) {
	p, exists := s.data[id]

	if !exists {
		return programmer{}, errors.New("programmer with current id does not exists")
	}

	return p, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func main() {

}
