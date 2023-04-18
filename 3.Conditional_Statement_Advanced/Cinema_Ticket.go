// Programm to print price of a thicket upon weekday, if:
// Monday Tuesday Wednesday Thursday Friday Saturday Sunday
// 12		12		14			14		12		16		16

package main

import "fmt"

func main() {

	var day string
	fmt.Scanln(&day)

	var price = map[string]int {

		"Monday": 12,
		"Tuesday": 12,
		"Wednesday": 14,
		"Thursday": 14,
		"Friday": 12,
		"Saturday": 16,
		"Sunday": 16,

	}

	value, exists := price[day]

	if exists { fmt.Println(value) }

}
