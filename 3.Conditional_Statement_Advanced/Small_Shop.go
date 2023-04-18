// Write a programm, which reads a product(string), town(string) and quantity(float).
// The programm estimate the price of the product in the town if:
// Sofia - coffee = 0.50, water = 0.80, beer = 1.20, sweets = 1.45, peanuts = 1.60
// Plovdiv - 0.40, 0.70, 1.15, 1.30, 1.50
// Varna - 0.45, 0.70, 1.10, 1.35, 1.55


package main

import "fmt"

func main() {

	var product, town string
	fmt.Scanln(&product)
	fmt.Scanln(&town)
	var quantity float32
	fmt.Scanln(&quantity)

	var allSum float32

	if town == "Sofia" {

		if product == "coffee" {

			allSum += quantity * 0.50

		} else if product == "water" {

			allSum += quantity * 0.80

		} else if product == "beer" {

			allSum += quantity * 1.20

		} else if product == "sweets" {

			allSum += quantity * 1.45

		} else if product == "peanuts" {

			allSum += quantity * 1.60

		}

	} else if town == "Plovdiv" {

		if product == "coffee" {

			allSum += quantity * 0.40

		} else if product == "water" {

			allSum += quantity * 0.70

		} else if product == "beer" {

			allSum += quantity * 1.15

		} else if product == "sweets" {

			allSum += quantity * 1.30

		} else if product == "peanuts" {

			allSum += quantity * 1.50

		}

	} else if town == "Varna" {

		if product == "coffee" {

			allSum += quantity * 0.45

		} else if product == "water" {

			allSum += quantity * 0.70

		} else if product == "beer" {

			allSum += quantity * 1.10

		} else if product == "sweets" {

			allSum += quantity * 1.35

		} else if product == "peanuts" {

			allSum += quantity * 1.55

		}
	
	}

	fmt.Println(allSum)

}
