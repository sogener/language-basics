package main

import "fmt"

type employee struct {
	name   string
	sex    string // пол
	age    int
	salary int // зарплата
}

func main() {
	firstEmployee := newEmployee("Вася", "М", 25, 1500)

	firstEmployee.setName("имя для struct через receiver")
	fmt.Println(firstEmployee.getInfo())
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}
func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d", e.name, e.sex, e.age, e.salary)
}
func (e *employee) setName(name string) {
	e.name = name
}
