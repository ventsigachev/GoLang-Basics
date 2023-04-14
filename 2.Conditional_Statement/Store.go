// You have a child store, where you selling toys:
// Teddy Bear - 4.10
// Puzzle - 2.60
// Doll - 3.00
// Truck - 2.00
// Minion - 8.20
// If purchased toys are >= 50 you give a discount of 25% of the all sum. You need to pay 10% of the sum for shop rent. 
// Inputs:
//		f_line - price of vacation
//		s_line - count of puzzles
//		t_line - count of dolls
//		fth_line - count of bears
//		fifth_line - count of minions
//		s_line - count of trucks
// Estimate if you can go on vacation.

package main

import (

	"fmt"
	"math"

)

func main() {

	var priceVacation float32
	fmt.Scanln(&priceVacation)
	var puzzles float32
	fmt.Scanln(&puzzles)
	var dolls float32
	fmt.Scanln(&dolls)
	var bears float32
	fmt.Scanln(&bears)
	var minions float32
	fmt.Scanln(&minions)
	var trucks float32
	fmt.Scanln(&trucks)

	var amount float32 = (puzzles * 2.60) + (dolls * 3.00) + (bears * 4.10) + (minions * 8.20) + (trucks * 2.00)
	var toysCount float32 = puzzles + dolls + bears + minions + trucks


	if toysCount >= 50 {

		var discount float32 = 0.25 * amount
		amount -= discount

	}

	var rent = 0.1 * amount
	amount -= rent

	var rest float64 = math.Abs(float64(amount - priceVacation))

	if amount >= priceVacation {

		fmt.Printf("Yes! %.2f lv left.", rest)

	} else {

		fmt.Printf("Not enough money! %.2f lv needed.", rest)
	}

}
