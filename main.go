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

	fmt.Printf("%+v\n", firstEmployee.name)
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
