// Write a programm which reads an int and prints the result:
// If the int is less than 100, the output is "Less than 100"
// If the int is between 100 - 200 "Between 100 and 200"
// If the int is greater than 200 "Greater than 200"

package main

import "fmt"

func main() {

	var i int
	fmt.Scanln(&i)

	if i < 100 {

		fmt.Println("Less than 100")

	} else if i <= 200 {

		fmt.Println("Between 100 and 200")

	}else {

		fmt.Println("Greater than 200")

	}
}
