package main

import "fmt"

func main() {
	var f_name string
	fmt.Scanln(&f_name)

	var l_name string
	fmt.Scanln(&l_name)

	var age int
	fmt.Scanln(&age)
	
	var town string
	fmt.Scanln(&town)

	fmt.Printf("You are %s %s, a %d-years old person from %s.", f_name, l_name, age, town)
}
