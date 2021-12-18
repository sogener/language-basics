package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(1)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)

	fmt.Printf("cicrleArea:\n", circleArea)
	fmt.Printf("err:\n", err)
	return

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr2\n")

	fmt.Printf("Площадь круга: %f32 см. кв.\n\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}

	return float32(radius) * float32(radius) * pi, nil
}
