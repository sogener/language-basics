package main

func main() {

	const currentMoney = 100

	setDeposit(currentMoney, true)
}

func setDeposit(money int, totalHold bool) string {
	if money < 100 {
		return "You need at least 100$ to byu something"
	}

	if totalHold {
		money := incrementValue(money)
		println(money)
	}

	return "Your deposit still the same"
}

func incrementValue(value int) int {
	return value * 2
}
