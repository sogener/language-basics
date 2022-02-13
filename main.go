package main

import "fmt"

func main() {
	type Employee struct {
		name   string
		sex    string
		age    int
		salary int
	}

	firstEmployee := Employee{
		name:   "Vasya",
		sex:    "Male",
		age:    18,
		salary: 1000,
	}
	secondEmployee := Employee{
		name:   "Alex",
		sex:    "Male",
		age:    22,
		salary: 555,
	}

	fmt.Printf("%+v\n", firstEmployee)
	fmt.Printf("%+v\n", secondEmployee)
}
