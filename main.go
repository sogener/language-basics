package main

import "fmt"

func main() {
	employee := struct {
		name   string
		sex    string
		age    int
		salary int
	}{
		name:   "Вася",
		sex:    "Male",
		age:    18,
		salary: 1000,
	}

	fmt.Printf("%+v\n", employee)
}
