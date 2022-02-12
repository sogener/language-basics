package main

import "fmt"

func main() {
	bankOptions := [4]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
		"дописать третий модуль",
	}

	// Мы создали пустой срез. Когда он не ссылается ни на какой массив, его значение по умолчанию равно nil.
	var actions []string
	fmt.Println(actions == nil)

	/*
		Далее мы присвоили ему срез нашего массива с 1 по 4 элемент.
		Синтаксис [m:n] позволяет делать выборку элементов массива или среза от индекса m до n.
		Пустые скобки [:] означают выборку всех элементов.
	*/
	actions = bankOptions[1:4]
	fmt.Println(actions == nil)

	for i := range actions {
		fmt.Println(actions[i])
	}
}
