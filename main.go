package main

import "fmt"

func main() {
	users := map[string]int{
		"user1": 1,
		"user2": 2,
	}

	users["Billy"] = 10
	users["Alex"] = 20

	delete(users, "user2")
	fmt.Println(users)
}
