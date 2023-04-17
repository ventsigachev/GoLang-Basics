// With Switch / Case

package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	switch num {

		case 1: fmt.Println("Monday")

		case 2: fmt.Print("Tuesday")

		case 3: fmt.Print("Wednesday")

		case 4: fmt.Print("Thursday")

		case 5: fmt.Print("Friday")

		case 6: fmt.Print("Saturday")

		case 7: fmt.Print("Sunday")

		default: fmt.Print("Error")


	}
}
