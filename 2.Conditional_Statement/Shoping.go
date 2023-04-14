// You need to buy video cards(vc), processors(p), ram(r).
// If vc > p - discount of 15%.
// Price of one vc = 250, price of processor = 35% of price of all vc, price of 1 ram = 10% of price of all vc.
// Inputs:
//		1. Budget, 2. VC, 3. P, 4. R


package main

import (

	"fmt"
	"math"

)

func main() {

	var b, vc, p, r float64
	fmt.Scanln(&b)
	fmt.Scanln(&vc)
	fmt.Scanln(&p)
	fmt.Scanln(&r)

	var priceVC float64 = vc * 250
	var pricePerProcessors float64 = priceVC * 0.35 * p
	var pricePerRams float64 = priceVC * 0.1 * r 
	var allSum float64 = pricePerProcessors + pricePerRams + priceVC

	if vc > p {

		allSum -= allSum * 0.15

	}

	var rest float64 = math.Abs(b - allSum)

	if b >= allSum {

		fmt.Printf("You have %.2f leva left!", rest)

	} else {

		fmt.Printf("Not enough money! You need %.2f leva more!", rest)

	}
}
