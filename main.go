package main

import "fmt"

func main() {
	users := map[string]int{
		"user1": 1,
		"user2": 2,
	}

	users["Billy"] = 10
	users["Alex"] = 20

	fmt.Println(users)

	// Init map with function make()
	var age = make(map[string]int)
	age["test11"] = 100

	fmt.Println(age)

	// Maps like [key => value] arrays
	for key, value := range users {
		fmt.Println(key, value)
	}

}
