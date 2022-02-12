package main

import "fmt"

func main() {

	/*
		Когда же мы выполнили функцию append(), то:
		Новый элемент вышел за рамки длины массива
		Создался новый массив, с длинною вдвое больше (4 -> 8)
		Срез начал ссылаться на новый массив
		Все значения из старого массива скопировались в новый + добавился еще один элемент
	*/

	var bankOptions = []string{
		"Просмотр баланса",
		"Удвоить баланс",
		"Уменьшить баланс",
		"Внести депозит",
	}
	fmt.Printf("Длинна списка %d\n", len(bankOptions))
	fmt.Printf("Емкость списка %d\n", cap(bankOptions))

	bankOptions = append(bankOptions, "Забрать депозит")

	fmt.Printf("Длинна списка %d\n", len(bankOptions))
	fmt.Printf("Емкость списка %d\n", cap(bankOptions))

	newBankOptions := append(bankOptions, "Забрать депозит")
	fmt.Printf("Длинна списка %d\n", len(newBankOptions))
	fmt.Printf("Емкость списка %d\n", cap(newBankOptions))
}
