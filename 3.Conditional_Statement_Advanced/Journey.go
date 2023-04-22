// Inputs:
//	1.Budget - float.
//	2.Season = Summer or Winter.
//Outputs:
//	1."Somewhere in {destination}" ,choose from "Bulgaria", "Balkans" or "Europe".
//	2."{Type of vacation} – {Money spent}".
//Constrictions:
//	Budget defines the destination and season - money spent(ms):
//	If budget <= 100, destination is Bulgaria and ms is 30% of budget if summer and 70% if winter.
//	If budget <= 1000, destination is Balkans, and ms is 40% - summer, 80% - winter.
//	If budget > 1000, destination is Europe, ms is 90% both seasons.



package main

import "fmt"

func main() {
	var budget float64
	fmt.Scanln(&budget)

	var season string
	fmt.Scanln(&season)

	var destination string = ""
	var price float64 = 0
	var place string = ""

	if budget <= 100 {

		destination = "Bulgaria"
		switch season {

		case "summer":
			price = budget * 0.3
			place = "Camp"
			break

		case "winter":
			price = budget * 0.7
			place = "Hotel"
			break

		}

	} else if budget <= 1000 {
		destination = "Balkans"
		switch season {

		case "summer":
			price = budget * 0.4
			place = "Camp"
			break

		case "winter":
			price = budget * 0.8
			place = "Hotel"
			break

		}

	} else if budget > 1000 {
		destination = "Europe"
		price = budget * 0.9
		place = "Hotel"
	}

	fmt.Printf("Somewhere in %s\n", destination)
	fmt.Printf("%s - %.2f", place, price)

}
