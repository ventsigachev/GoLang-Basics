// Programm to choose outfit, based upon weather and time of the day.
//	Inputs: 1. Temperature(int), 2. Time of the day: Morning, Afternoon, Evening.
//	Мorning - 10 <= grades <= 18 Outfit = Sweatshirt Shoes = Sneakers
//  Afternoon - Outfit = Shirt, Shoes = Moccasins
//  Evening - Outfit = Shirt, Shoes = Moccasins
//  18 < grades <= 24: Moring - Outfit­ = Shirt Shoes = Moccasins
//  Afernoon - Outfit = T-Shirt Shoes = Sandals
//  Evening - Outfit = Shirt Shoes = Moccasins
//  grades >= 25: Morning - Outfit = T-Shirt Shoes = Sandals
//	Afternoon - Outfit = Swim Suit Shoes = Barefoot
//	Evening - Outfit = Shirt Shoes = Moccasins

package main

import "fmt"

func main () {
	var degrees int 
	fmt.Scanln(&degrees)

	var partOfDay string
	fmt.Scanln(&partOfDay)

	var outfit string = ""
	var shoes string = "" 

	if degrees >= 10 && degrees <= 18 {
		
		switch partOfDay {
			case "Morning":
				outfit = "Sweatshirt"
				shoes = "Sneakers"
			case "Afternoon", "Evening":
				outfit = "Shirt"
				shoes = "Moccasins"
		}
	} else if degrees > 18 && degrees <= 24 {

		switch partOfDay {
			case "Morning", "Evening":
				outfit = "Shirt"
				shoes = "Moccasins"
			case "Afternoon":
				outfit = "T-Shirt"
				shoes = "Sandals"
		}

	} else if degrees >= 25 {

		switch partOfDay {
			case "Morning":
				outfit = "T-Shirt"
				shoes = "Sandals"
			case "Afternoon":
				outfit = "Swim Suit"
				shoes = "Barefoot"
			case "Evening":
				outfit = "Shirt"
				shoes = "Moccasins"
		}

	}

	fmt.Printf("It's %d degrees, get your %s and %s.", degrees, outfit, shoes)

}
