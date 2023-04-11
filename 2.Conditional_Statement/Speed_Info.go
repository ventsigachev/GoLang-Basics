// Input speed - real number
// Output:
//		speed <= 10: "slow"
//		between 10 and 50: "average"
//		above 50 below/equal 150:"fast"
//		above 150 below/equal 1000: "ultra fast"
//		above 1000: "extremely fast"

package main

import "fmt"

func main() {

	var speed float32
	fmt.Scanln(&speed)

	if speed <= 10 {

		fmt.Println("slow")

	} else if speed <= 50 {

		fmt.Println("average")

	} else if speed <= 150 {

		fmt.Println("fast")

	} else if speed <= 1000 {

		fmt.Println("ultra fast")

	} else {

		fmt.Println("extremely fast")

	}

}
