// Input is an integer - a start points. Write a programm to calculate bonuses, taking in advise the following:
// If input <= 100 bonus is 5 points.
// If input > 100, bonus is 20% of initial points.
// If input > 1000, bonus is 10%.
// If input is even number, bonus is 1 point.
// If input number ends with 5, bonus is 2 points.

package main

import "fmt"

func main() {

	var number int
	fmt.Scanln(&number)

	var n float32

	if number <= 100 {

		n = 5

	} else if number > 1000 {

		n = float32(number) * 0.1
		
	} else {
		
		n = float32(number) * 0.2
	}
	
	if number % 2 == 0 {

		n += 1

	} else if number % 10 == 5 {

		n += 2

	}

	fmt.Println(n)
	fmt.Println(float32(number) + n)

}
