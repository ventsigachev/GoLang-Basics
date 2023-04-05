// Write a programm to estimate the cost of landscaping. The price for square meter is 7.61.
// You have a discount of 18 %.
// Inputs:
//		One line - sq.m. - real number[0.00 .... 10000.00]
// Output:
//		First line - the price for landscaping
//		Second line - the discount


package main

import "fmt"

func main() {
	var area float64
	fmt.Scanln(&area)

	var generalPrice float64 = area * 7.61
	var discount float64 = generalPrice * 0.18
	var endPrice float64 = generalPrice - discount

	fmt.Printf("The final price for landscaping is %.2f \n", endPrice)
	fmt.Printf("The discount is %.2f", discount)
}