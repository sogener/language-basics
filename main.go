package main

import "fmt"

func main() {
	bankOptions := [4]string{
		"Пополнить",
		"Просмотреть",
		"Снять",
		"Посмотреть %",
	}

	// Мы создали пустой срез. Когда он не ссылается ни на какой массив, его значение по умолчанию равно nil.
	var actions []string
	/*
		Далее мы присвоили ему срез нашего массива с 1 по 4 элемент.
		Синтаксис [m:n] позволяет делать выборку элементов массива или среза от индекса m до n.
		Пустые скобки [:] означают выборку всех элементов.
	*/
	actions = bankOptions[0:4]

	for i := range actions {
		fmt.Println(actions[i])
	}

	fmt.Println("---- После changeOptions ----")
	changeOptions(actions)

	for i := range actions {
		println(actions[i])
	}
}
func changeOptions(actions []string) {
	actions[0] = "Изменили первый элем1"
	actions[1] = "Изменили 2"
}
