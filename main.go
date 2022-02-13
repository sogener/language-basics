package main

import "fmt"

func main() {
	type employee struct {
		name   string
		sex    string // пол
		age    int
		salary int // зарплата
	}

	firstEmployee := newEmployee("Вася", "М", 25, 1500)

	fmt.Printf("%+v\n", firstEmployee)
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}
