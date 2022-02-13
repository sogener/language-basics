package main

import (
	"errors"
	"fmt"
)

type programmer struct {
	id        int
	name      string
	isWorking bool
}
type dumbStorage struct{}

type storage interface {
	insert(p programmer) error
	get(id int) (programmer, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]programmer
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]programmer),
	}
}

func (s *memoryStorage) insert(e programmer) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (programmer, error) {
	e, exists := s.data[id]
	if !exists {
		return programmer{}, errors.New("programmer with such id doesn't exist")
	}

	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s dumbStorage) insert(p programmer) error {
	fmt.Println("Вставка прошла успешно")
	return nil
}

func (s dumbStorage) get(id int) (programmer, error) {
	p := programmer{id: id}
	return p, nil
}

func (s dumbStorage) delete(id int) error {
	fmt.Println("Удаление произошло")
	return nil
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(programmer{id: i})
	}
}

func main() {
	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)
}
