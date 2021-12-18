package main

import "fmt"

func main() {
	x := 1
	y := 2
	numberSwapper(&x, &y)

	fmt.Println(x)
	fmt.Println(y)

}
func incrementNumber(number *int) {
	*number += 1
}
func numberSwapper(firstNumber, secondNumber *int) {

	// Default
	holder := *firstNumber
	*firstNumber = *secondNumber
	*secondNumber = holder

	// Go
	*firstNumber, *secondNumber = *secondNumber, *firstNumber

}
