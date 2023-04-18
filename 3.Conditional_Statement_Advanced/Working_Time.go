// Programm to represent working time of an office(open/close), if day working time 10 - 18, from Monday to Saturday

package main

import "fmt"

func main() {

	var hour int
	fmt.Scanln(&hour)
	var day string
	fmt.Scanln(&day)

	if day == "Sunday" || hour < 10 || hour > 18 {

		fmt.Println("closed")

	} else {

		fmt.Println("open")

	}

}
