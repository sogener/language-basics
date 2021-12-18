package main

import "fmt"

func main() {
	x := 15

	fmt.Println("значение x до изменения значения: ", x)
	incrementNumber(&x)
	fmt.Println("значение x после: ", x)

	z := new(int)
	incrementNumber(z)
	fmt.Println("адрес памяти z, созданной с помощью new() -  ", z)
	fmt.Println("значение z, созданной с помощью new() -  ", *z)

	//	Пустой указатель, который вернёт nil
	var p *int
	fmt.Println(p)
}
func incrementNumber(number *int) {
	*number += 1
}
