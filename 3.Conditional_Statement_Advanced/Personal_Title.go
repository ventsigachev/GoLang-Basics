// Write a programm which reads age(float) and a gender(string), and prints personal title.
//  • "Mr." - gender male(m), age >= 16 
//	• "Master" - gender male(m), age < 16
//	• "Ms." - gender woman(f), age >= 16
//	• "Miss" - gender woman(f), age < 16

package main

import "fmt"

func main() {

	var age float32
	fmt.Scanln(&age)

	var gender string
	fmt.Scanln(&gender)

	if gender == "m" {
		
		switch {
			case age < 16:

				fmt.Println("Master")

			case age >= 16:

				fmt.Println("Mr.")

		}
	} else if gender == "f" {
		
		if age < 16 {

			fmt.Println("Miss")

		} else {
			
			fmt.Println("Ms.")
		}
	}

}
