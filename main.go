package main

import "fmt"

func main() {
	users := map[string]int{
		"user1": 1,
		"user2": 2,
	}

	users["Billy"] = 10
	users["Alex"] = 20

	/*
		Как мы уже говорили, мапы, как и срезы, являются указателями на область в памяти.
		Поэтому если скопировать мапу в новую переменную и удалить из нее элементы,
		это отобразится также и на новой мапе.
	*/

	var usersNew map[string]int
	usersNew = users

	delete(usersNew, "user2")
	fmt.Println(users)

}
