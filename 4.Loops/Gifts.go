package main

import (
	"fmt"
	"math"
)

func main() {
	var age, priceToy int
	var priceMachine float64
	fmt.Scanln(&age)
	fmt.Scanln(&priceMachine)
	fmt.Scanln(&priceToy)

	var savedMoney float64 = 0
	var countToys int = 0

	for birthday := 1; birthday <= age; birthday++ {
		if birthday % 2 == 0 {

			savedMoney += float64(birthday) * 5
			savedMoney--

		} else {

			countToys++

		}
	}
	
	savedMoney += (float64(countToys) * float64(priceToy))

	var diff float64 = math.Abs(savedMoney - priceMachine)
	if savedMoney >= priceMachine {

		fmt.Printf("Yes! %.2f", diff)

	} else {

		fmt.Printf("No! %.2f", diff)

	}
}
