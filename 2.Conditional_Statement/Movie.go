// You are hired to write a programm to estimate the budget of a movie.
// The decor is 10% of budget.
// If stunt-men are more than 150, discount for clothes is 10%
// Inputs:
//		f_line - budget
//		s_line - count of stunts
//		t_line - price of clothes

package main

import (

	"fmt"
	"math"

)

func main() {

	var budget, stuntsCount, clothePrice float64
	fmt.Scanln(&budget)
	fmt.Scanln(&stuntsCount)
	fmt.Scanln(&clothePrice)

	var decor float64 = budget * 0.1
	var clothesPrice float64 = stuntsCount * clothePrice

	if stuntsCount > 150 {

		clothesPrice = clothesPrice - clothesPrice * 0.1
	}

	var endSum float64 = decor + clothesPrice

	var rest float64 = math.Abs(budget - endSum)

	if endSum <= budget {

		fmt.Printf("Action!\nWingard starts filming with %.2f leva left.", rest)

	} else {

		fmt.Printf("Not enough money!\nWingard needs %.2f leva more.", rest)
	}

}
