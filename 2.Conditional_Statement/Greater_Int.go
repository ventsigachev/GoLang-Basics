// Write a programm, which prints the greater of two digits.
// If the digits are equal, print also the digit.

package main

import "fmt"

func main() {

	var a int
	fmt.Scanln(&a)
	var b int
	fmt.Scanln(&b)

	if a > b {

		fmt.Println(a)

	} else if b > a {

		fmt.Println(b)

	} else {

		fmt.Println(a)
	}

}
