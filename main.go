package main

import "fmt"

func main() {
	users := map[string]int{
		"user1": 1,
		"user2": 2,
	}

	users["Billy"] = 10
	users["Alex"] = 20

	_, exists := users["Billy_not_found"]

	if !exists {
		fmt.Println("Billy нет в списке")
		return
	}

	fmt.Printf("Значение билли равно %d\n", users["Billy"])
}
