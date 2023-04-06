//	Deposit = deposit_cash + deposit_term * ((deposit_cash * annual_interest) / 12)
// Inputs:
//		First line - deposit_cash(dc)
//		Second line - deposit_term(dt) in mounts
//		Third line - annual_interest(ai)


package main

import "fmt"

func main() {
	
	var dc float32
	fmt.Scanln(&dc)
	var dt float32
	fmt.Scanln(&dt)
	var ai float32
	fmt.Scanln(&ai)

	var monthInterest = (dc * ai / 100) / 12
	var finalCash = dc + dt * monthInterest
	
	fmt.Println(finalCash)

}
