package main

import "fmt"

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

func main() {
	var s storage
	fmt.Println("s: ", s)
	fmt.Printf("type of s: %T\n", s)

	s = newDumbStorage()
	fmt.Println("s: ", s)
	fmt.Printf("type of s: %T\n", s)
}
