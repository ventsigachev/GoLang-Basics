// Write a program to estimate necessities of food for dogs and cats.
// One package of dog food costs 2.50 and cat food is 4.
// Input:
//		First line - packages of dog food
//		Second line - packages of cat food 


package main

import "fmt"

func main() {

	var dogPackages int
	fmt.Scanln(&dogPackages)

	var catPackages int
	fmt.Scanln(&catPackages)

	var dogCost float64 = float64(dogPackages) * 2.50
	var catCost float64 = float64(catPackages) * 4

	var allCost = dogCost + catCost

	fmt.Println(allCost)

}
