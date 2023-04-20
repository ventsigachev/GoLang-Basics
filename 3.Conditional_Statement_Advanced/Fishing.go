package main

import "fmt"


func main () {
	

	var budget int
	fmt.Scanln(&budget)

	var season string
	fmt.Scanln(&season) //"Spring", "Summer", "Autumn", "Winter"

	var countFishers int 
	fmt.Scanln(&countFishers)

	//	boat rent -> depends on season
	var rent float32 = 0
	switch season {
		case "Spring":
			rent = 3000
		case "Autumn", "Summer":
			rent = 4200
		case "Winter":
			rent = 2600

	}

	//	discount -> depends on fishmans counts
	if countFishers <= 6 {
		
		rent = rent - 0.10 * rent 
	} else if countFishers >= 7 && countFishers <= 11 {
		
		rent = rent - 0.15 * rent //0.85 * rent
	} else if countFishers >= 12 {

		rent = rent - 0.25 * rent
	}

	
	if countFishers % 2 == 0 && season != "Autumn" {
	
		rent = rent - 0.05 * rent // 0.95 * rent
	}


	if float32(budget) >= rent {
		var leftMoney float32 = float32(budget) - rent
		fmt.Printf("Yes! You have %.2f leva left.", leftMoney)
	} else {
		//не е достатъчен
		var needMoney float32 = rent - float32(budget)
		fmt.Printf("Not enough money! You need %.2f leva.", needMoney)
	}

}
