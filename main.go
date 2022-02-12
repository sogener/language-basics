package main

import (
	"errors"
	"fmt"
)

func main() {

	const currentMoney = 101

	setDeposit(currentMoney, true)
}

func setDeposit(money int32, isHold bool) int32 {
	money, err := incrementValue(money, isHold)

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	fmt.Println(money)
	return money
}

func incrementValue(value int32, isHold bool) (int32, error) {
	if value < 100 {
		return int32(0), errors.New("you need at least 100$ to byu something")
	}

	if !isHold {
		return value, nil
	}

	return value * 2, nil
}
