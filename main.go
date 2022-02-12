package main

import (
	"fmt"
)

func main() {
	const currentMoney = 101

	var bankOptions = [3]string{
		"Просмотреть баланс",
		"Увеличить x2",
		"Уменьшить x2",
	}

	var todoTest = [5]string{"Не обязательно заполнять весь массив"}

	fmt.Printf("Кол-во элементов в массиве: %d\n", len(bankOptions))
	fmt.Printf("Кол-во элементов в массиве: %d\n", len(todoTest))

	for key, value := range bankOptions {
		fmt.Printf("%d. %s\n", key, value)
	}
}
