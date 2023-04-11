// Write s programm which estimate a grade, readed from a console.
// If the grade is equal or greater than 5, the output is excellent.

package main

import "fmt"

func main() {
	var grade int
	fmt.Scanln(&grade)

	if grade >= 5 {

		fmt.Println("Excellent!")
		
	}
}
