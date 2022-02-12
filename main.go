package main

import "fmt"

func main() {
	var arr [3]int
	pArr := &arr
	fillArray(pArr)

	fmt.Println(arr)

}
func fillArray(pArr *[3]int) {
	for i := 0; i < len(pArr); i++ {
		pArr[i] = i
	}
	fmt.Println("Fill array: ", pArr)
}
