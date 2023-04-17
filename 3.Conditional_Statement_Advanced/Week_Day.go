// Programm to display day of the week by user input(1 to 7), or ERROR.

package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	var week = map[int]string {

		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	value, exist := week[num]

	if exist {

		fmt.Println(value)

	} else {

		fmt.Println("Error")

	}

}
