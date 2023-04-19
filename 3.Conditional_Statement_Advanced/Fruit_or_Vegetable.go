// Programm to define fruit or vegetable, by user input.

package main

import "fmt"

func main() {

	var i string
	fmt.Scanln(&i)

	switch i {

	case "banana", "apple", "kiwi", "cherry", "lemon", "grapes":
		fmt.Println("fruit")
	
	case "tomato", "cucumber", "pepper", "carrot":
		fmt.Println("vegatable")

	default:
		fmt.Println("uknown")

	}

}
